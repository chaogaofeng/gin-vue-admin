package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/upay"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strconv"
	"time"
)

//@function: ScanTronUSDT
//@description: 扫描 USDT 交易
//@param: db(数据库对象) *gorm.DB, tableName(表名) string, compareField(比较字段) string, interval(间隔) string
//@return: error

func ScanTronUSDT(db *gorm.DB) error {
	// 查找所有过期的订单并更新状态
	currentTime := time.Now()
	if err := db.Model(&upay.PayOrder{}).
		Where("status = ? AND expired_at < ?", 0, currentTime).
		Update("status", "2").Error; err != nil {
		return fmt.Errorf("update order status=2 error: %v", err)
	}

	var orders []upay.PayOrder
	query := `
		SELECT DISTINCT receiver
		FROM upay_pay_order
		WHERE status = "0";
	`
	err := db.Raw(query).Scan(&orders).Error
	if err != nil {
		return fmt.Errorf("fetching order status=0 error: %v", err)
	}

	baseUrl := global.GVA_CONFIG.TronGrid.BaseURL
	apiKeys := global.GVA_CONFIG.TronGrid.ApiKeys
	contractAddress := global.GVA_CONFIG.TronGrid.TRC20
	if len(baseUrl) == 0 || len(apiKeys) == 0 || len(contractAddress) == 0 {
		return nil
	}

	for _, order := range orders {
		address := order.Receiver
		for _, apiKey := range apiKeys {
			transactions, err := findTRC20(baseUrl, apiKey, address, order.CreatedAt.Unix(), contractAddress)
			if err != nil {
				return fmt.Errorf("fetching order tx error: %v", err)
			}

			for _, transaction := range transactions {
				value, _ := decimal.NewFromString(transaction.Value)
				value = value.Div(decimal.NewFromFloat(10).Mul(decimal.NewFromInt(int64(transaction.TokenInfo.Decimals))))
				completedAt := time.Now()
				err := db.Model(&upay.PayOrder{}).
					Where("amount = ?", value).
					Updates(upay.PayOrder{
						Status:       "1",
						CompletedAt:  &completedAt,
						Hash:         transaction.TransactionID,
						ActualAmount: value,
						NotifyAt:     &completedAt,
					}).Error
				if err != nil {
					return fmt.Errorf("update order status=1 error: %v", err)
				}
				if len(order.NotifyUrl) == 0 {
					continue
				}
				var updatedOrder upay.PayOrder
				if err := db.Model(&upay.PayOrder{}).
					Where("amount = ?", value).
					First(&updatedOrder).
					Error; err == nil {
					var app upay.APP
					if err = global.GVA_DB.Where("app_id = ?", order.AppID).First(&app).Error; err == nil {
						req := map[string]string{
							"appId":           order.AppID,
							"orderNo":         order.OrderNo,
							"merchantOrderNo": order.MerchantOrderNo,
							"chainType":       order.ChainType,
							"crypto":          order.Amount.String(),
							"actualCrypto":    order.ActualAmount.String(),
							"poundage":        order.Fee.String(),
							"actualPoundage":  order.ActualFee.String(),
							"status":          order.Status,
							"createdAt":       strconv.FormatInt(order.CreatedAt.Unix(), 10),
							"completedAt": strconv.FormatInt(func() int64 {
								if order.CompletedAt.IsZero() {
									return 0 // 返回默认值
								}
								return order.CompletedAt.Unix()
							}(), 10),
						}
						expectedSignature := utils.GenerateSignature(req, app.AppSecret)
						req["signature"] = expectedSignature
						data, _ := json.Marshal(req)
						// 通知
						resp, err := http.Post(order.NotifyUrl, "application/json", bytes.NewBuffer(data))
						if err != nil {
						}
						defer resp.Body.Close()
					}
				}
			}
		}
	}
	return nil
}

// TransactionTRC20 data structure
type TransactionTRC20 struct {
	TransactionID  string `json:"transaction_id"`
	BlockTimestamp int64  `json:"block_timestamp"`
	From           string `json:"from"`
	To             string `json:"to"`
	Value          string `json:"value"`
	Type           string `json:"type"`
	TokenInfo      struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
		Address  string `json:"address"`
	} `json:"token_info"`
}

// ResponseTRC20 structure
type ResponseTRC20 struct {
	Data    []TransactionTRC20 `json:"data"`
	Success bool               `json:"success"`
	Meta    struct {
		Links struct {
			Next string `json:"next"`
		} `json:"links"`
	} `json:"meta"`
}

func findTRC20(baseUrl string, apiKey string, address string, minTimestamp int64, contractAddress string) ([]TransactionTRC20, error) {
	onlyTo := "true"
	onlyConfirmed := "true"
	limit := 200

	url := fmt.Sprintf("%s/v1/accounts/%s/transactions/trc20?only_confirmed=%s&only_to=%s&limit=%d&contract_address=%s&min_timestamp=%d",
		baseUrl, address, onlyConfirmed, onlyTo, limit, contractAddress, minTimestamp)

	var txs []TransactionTRC20
	for {
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("accept", "application/json")
		if len(apiKey) > 0 {
			req.Header.Add("TRON-PRO-API-KEY", apiKey)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("http error %v\n", err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return nil, fmt.Errorf("status code %d %s\n", res.StatusCode, res.Status)
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("read body error: %v", err)
		}

		var apiResponse ResponseTRC20
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			return nil, fmt.Errorf("parsing json error: %v", err)
		}

		// Process transactions
		for _, tx := range apiResponse.Data {
			fmt.Printf("Transaction ID: %s\nFrom: %s\nTo: %s\nValue: %s\nTimestamp: %d\n\n",
				tx.TransactionID, tx.From, tx.To, tx.Value, tx.BlockTimestamp)
		}
		txs = append(txs, apiResponse.Data...)

		// Check if there's a next link for pagination
		if apiResponse.Meta.Links.Next == "" {
			break
		}
		// Update URL for the next page
		url = apiResponse.Meta.Links.Next
		fmt.Println(string(body))
	}

	return nil, nil
}
