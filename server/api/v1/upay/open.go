package upay

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/upay"
	upayReq "github.com/flipped-aurora/gin-vue-admin/server/model/upay/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
)

const (
	ERROR   = 0
	SUCCESS = 1
)

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

type OpenApi struct{}

func (openApi *OpenApi) OrderApply(c *gin.Context) {
	var req upayReq.OrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			fmt.Sprintf("invalid request: %v", err.Error()),
		})
		return
	}

	amount, err := decimal.NewFromString(req.FiatAmount)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			fmt.Sprintf("invalid fiatAmount: %v", err.Error()),
		})
		return
	}

	app, err := appService.GetAPPByAppID(req.AppID)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			fmt.Sprintf("invalid appId: %v", err.Error()),
		})
		return
	}

	// 将请求参数放入 map 中以生成签名
	params := map[string]string{
		"appId":           req.AppID,
		"merchantOrderNo": req.MerchantOrderNo,
		"chainType":       req.ChainType,
		"fiatAmount":      req.FiatAmount,
		"fiatCurrency":    req.FiatCurrency,
		"notifyUrl":       req.NotifyUrl,
	}

	_, total, err := walletAddressService.GetWalletAddressInfoList(upayReq.WalletAddressSearch{
		UserID:    app.UserID,
		ChainType: req.ChainType,
		Status:    "0",
		PageInfo: request.PageInfo{
			Page:     1,
			PageSize: 10,
		},
	})
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			fmt.Sprintf("Invaid address: %v", err.Error()),
		})
		return
	}
	if total == 0 {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			"Not found address",
		})
		return
	}

	// 生成签名
	expectedSignature := utils.GenerateSignature(params, app.AppSecret)

	// 校验签名
	if req.Signature != expectedSignature {
		c.JSON(http.StatusOK, Response{
			ERROR,
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
			ERROR,
			map[string]interface{}{},
			fmt.Sprintf("%v: %v", global.Translate("general.creationFailErr"), err.Error()),
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
		SUCCESS,
		ret,
		global.Translate("general.createSuccess"),
	})
}

func (openApi *OpenApi) OrderSearch(c *gin.Context) {
	var req upayReq.SearchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			fmt.Sprintf("invalid request: %v", err.Error()),
		})
		return
	}

	app, err := appService.GetAPPByAppID(req.AppID)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			fmt.Sprintf("invalid appId: %v", err.Error()),
		})
		return
	}

	// 将请求参数放入 map 中以生成签名
	params := map[string]string{
		"appId":           req.AppID,
		"merchantOrderNo": req.MerchantOrderNo,
	}

	// 生成签名
	expectedSignature := utils.GenerateSignature(params, app.AppSecret)

	// 校验签名
	if req.Signature != expectedSignature {
		c.JSON(http.StatusOK, Response{
			ERROR,
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
			ERROR,
			map[string]interface{}{},
			fmt.Sprintf("invalid appId or merchantOrderNo: %v", err.Error()),
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
		CreatedAt:       strconv.FormatInt(order.CreatedAt.Unix(), 10),
		CompletedAt: strconv.FormatInt(func() int64 {
			if order.CompletedAt.IsZero() {
				return 0 // 返回默认值
			}
			return order.CompletedAt.Unix()
		}(), 10),
	}

	c.JSON(http.StatusOK, Response{
		SUCCESS,
		ret,
		global.Translate("general.getDataSuccess"),
	})
}

func (openApi *OpenApi) QueryAll(c *gin.Context) {
	var req upayReq.QueryAllRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			fmt.Sprintf("invalid request: %v", err.Error()),
		})
		return
	}

	app, err := appService.GetAPPByAppID(req.AppID)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			0,
			map[string]interface{}{},
			fmt.Sprintf("invalid appId: %v", err.Error()),
		})
		return
	}

	params := map[string]string{
		"appId":    req.AppID,
		"pageNo":   req.PageNo,
		"pageSize": req.PageSize,
	}

	expectedSignature := utils.GenerateSignature(params, app.AppSecret)

	if req.Signature != expectedSignature {
		c.JSON(http.StatusOK, Response{
			ERROR,
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
			ERROR,
			map[string]interface{}{},
			fmt.Sprintf("invalid appId: %v", err.Error()),
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
			CreatedAt:       strconv.FormatInt(order.CreatedAt.Unix(), 10),
			CompletedAt: strconv.FormatInt(func() int64 {
				if order.CompletedAt.IsZero() {
					return 0 // 返回默认值
				}
				return order.CompletedAt.Unix() // 返回实际值
			}(), 10),
		})
	}

	ret := PageResult{
		Total:    strconv.Itoa(int(total)),
		Page:     req.PageNo,
		PageSize: req.PageSize,
		List:     orders,
	}
	c.JSON(http.StatusOK, Response{
		SUCCESS,
		ret,
		global.Translate("general.getDataSuccess"),
	})
}
