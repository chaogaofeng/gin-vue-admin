package upay

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/upay"
	upayReq "github.com/flipped-aurora/gin-vue-admin/server/model/upay/request"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

type OpenApi struct{}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    string      `json:"total"`
	Page     string      `json:"page"`
	PageSize string      `json:"pageSize"`
}

// 生成签名的函数
func generateSignature(params map[string]string, appSecret string) string {
	// 删除 signature 参数（因为不参与签名）
	delete(params, "signature")

	// 提取并按 ASCII 顺序排序参数名
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接成 key=value&key2=value2… 格式的字符串
	var stringA strings.Builder
	for i, k := range keys {
		stringA.WriteString(fmt.Sprintf("%s=%s", k, url.QueryEscape(params[k])))
		if i < len(keys)-1 {
			stringA.WriteString("&")
		}
	}

	// 在 stringA 末尾添加 appSecret
	stringSignTemp := stringA.String() + "&appSecret=" + appSecret

	// 进行 MD5 运算
	hash := md5.New()
	hash.Write([]byte(stringSignTemp))
	signature := strings.ToUpper(hex.EncodeToString(hash.Sum(nil))) // 转换为大写

	return signature
}

func (openApi *OpenApi) OrderApply(c *gin.Context) {
	var req upayReq.OrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	amount, err := decimal.NewFromString(req.FiatAmount)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	app, err := appService.GetAPPByAppID(req.AppID)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	appSecret := app.AppSecret

	// 将请求参数放入 map 中以生成签名
	params := map[string]string{
		"appId":           req.AppID,
		"merchantOrderNo": req.MerchantOrderNo,
		"chainType":       req.ChainType,
		"fiatAmount":      req.FiatAmount,
		"fiatCurrency":    req.FiatCurrency,
		"notifyUrl":       req.NotifyUrl,
	}

	// 生成签名
	expectedSignature := generateSignature(params, appSecret)

	// 校验签名
	if req.Signature != expectedSignature {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			"Invalid signature",
		})
		return
	}

	order := &upay.PayOrder{
		UserID:          app.UserID,
		AppID:           req.AppID,
		MerchantOrderNo: req.MerchantOrderNo,
		ChainType:       req.ChainType,
		Amount:          amount,
		Attach:          req.Attach,
		NotifyUrl:       req.NotifyUrl,
		RedirectUrl:     req.RedirectUrl,
		//req.ProductName,
		//req.FiatCurrency,
	}
	err = payOrderService.CreatePayOrder(order)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	payUrl := "https://example.com/pay"
	ret := upayReq.OrderResponse{
		AppID:           req.AppID,
		OrderNo:         order.OrderNo,
		MerchantOrderNo: req.MerchantOrderNo,
		ExchangeRate:    "1",
		Crypto:          order.Amount.String(),
		Status:          order.Status,
		PayUrl:          payUrl,
		Address:         order.Receiver,
		CreatedAt:       order.CreatedAt.Unix(),
		ExpiredAt:       order.ExpiredAt.Unix(),
	}

	c.JSON(http.StatusOK, Response{
		1,
		ret,
		"success",
	})
}

func (openApi *OpenApi) OrderSearch(c *gin.Context) {
	var req upayReq.SearchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	app, err := appService.GetAPPByAppID(req.AppID)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	appSecret := app.AppSecret

	// 将请求参数放入 map 中以生成签名
	params := map[string]string{
		"appId":           req.AppID,
		"merchantOrderNo": req.MerchantOrderNo,
	}

	// 生成签名
	expectedSignature := generateSignature(params, appSecret)

	// 校验签名
	if req.Signature != expectedSignature {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			"Invalid signature",
		})
		return
	}

	list, _, err := payOrderService.GetPayOrderInfoList(upayReq.PayOrderSearch{
		AppID:           app.AppId,
		MerchantOrderNo: req.MerchantOrderNo,
	})
	if err != nil {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}
	order := list[0]

	ret := upayReq.Order{
		AppID:           req.AppID,
		OrderNo:         order.OrderNo,
		MerchantOrderNo: order.MerchantOrderNo,
		ChainType:       order.ChainType,
		Crypto:          order.Amount.String(),
		ActualCrypto:    order.ActualAmount.String(),
		Poundage:        order.Fee.String(),
		ActualPoundage:  order.ActualFee.String(),
		Status:          order.Status,
		Attach:          order.Attach,
		CreatedAt:       order.CreatedAt.Unix(),
		CompletedAt: func() int64 {
			if order.CompletedAt.IsZero() {
				return 0 // 返回默认值
			}
			return order.CompletedAt.Unix() // 返回实际值
		}(),
	}

	c.JSON(http.StatusOK, Response{
		1,
		ret,
		"success",
	})
}

func (openApi *OpenApi) QueryAll(c *gin.Context) {
	var req upayReq.QueryAllRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	app, err := appService.GetAPPByAppID(req.AppID)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	appSecret := app.AppSecret

	params := map[string]string{
		"appId":    req.AppID,
		"pageNo":   req.PageNo,
		"pageSize": req.PageSize,
	}

	expectedSignature := generateSignature(params, appSecret)

	if req.Signature != expectedSignature {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			"Invalid signature",
		})
		return
	}

	pageNo, _ := strconv.Atoi(req.PageNo)
	pageSize, _ := strconv.Atoi(req.PageSize)

	list, total, err := payOrderService.GetPayOrderInfoList(upayReq.PayOrderSearch{
		AppID: app.AppId,
		PageInfo: request.PageInfo{
			Page:     pageNo,
			PageSize: pageSize,
		},
	})
	if err != nil {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	var orders []upayReq.Order
	for _, order := range list {
		orders = append(orders, upayReq.Order{
			AppID:           req.AppID,
			OrderNo:         order.OrderNo,
			MerchantOrderNo: order.MerchantOrderNo,
			ChainType:       order.ChainType,
			Crypto:          order.Amount.String(),
			ActualCrypto:    order.ActualAmount.String(),
			Poundage:        order.Fee.String(),
			ActualPoundage:  order.ActualFee.String(),
			Status:          order.Status,
			Attach:          order.Attach,
			CreatedAt:       order.CreatedAt.Unix(),
			CompletedAt: func() int64 {
				if order.CompletedAt.IsZero() {
					return 0 // 返回默认值
				}
				return order.CompletedAt.Unix() // 返回实际值
			}(),
		})
	}

	ret := PageResult{
		Total:    strconv.Itoa(int(total)),
		Page:     req.PageNo,
		PageSize: req.PageSize,
		List:     orders,
	}
	c.JSON(http.StatusOK, Response{
		1,
		ret,
		"success",
	})
}
