package request

type OrderRequest struct {
	AppID           string `json:"appId" binding:"required"`
	MerchantOrderNo string `json:"merchantOrderNo" binding:"required"`
	ChainType       string `json:"chainType" binding:"required"`
	FiatAmount      string `json:"fiatAmount" binding:"required"`
	FiatCurrency    string `json:"fiatCurrency" binding:"required"`
	Attach          string `json:"attach,omitempty"`
	ProductName     string `json:"productName,omitempty"`
	NotifyUrl       string `json:"notifyUrl" binding:"required"`
	RedirectUrl     string `json:"redirectUrl,omitempty"`
	Signature       string `json:"signature" binding:"required"`
}

type OrderResponse struct {
	AppID           string `json:"appId"`
	OrderNo         string `json:"orderNo"`
	MerchantOrderNo string `json:"merchantOrderNo"`
	ExchangeRate    string `json:"exchangeRate"`
	Crypto          string `json:"crypto"`
	Status          string `json:"status"`
	PayUrl          string `json:"payUrl"`
	Address         string `json:"address"`
	CreatedAt       int64  `json:"createdAt"`
	ExpiredAt       int64  `json:"expiredAt"`
}

type SearchRequest struct {
	AppID           string `json:"appId" binding:"required"`
	MerchantOrderNo string `json:"merchantOrderNo" binding:"required"`
	Signature       string `json:"signature" binding:"required"`
}

type QueryAllRequest struct {
	AppID     string `json:"appId" binding:"required"`
	PageNo    string `json:"pageNo" binding:"required"`
	PageSize  string `json:"pageSize" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type Order struct {
	AppID           string `json:"appId"`
	OrderNo         string `json:"orderNo"`
	MerchantOrderNo string `json:"merchantOrderNo"`
	ChainType       string `json:"chainType"`
	Crypto          string `json:"crypto"`
	ActualCrypto    string `json:"actualCrypto"`
	Poundage        string `json:"poundage"`
	ActualPoundage  string `json:"actualPoundage"`
	Status          string `json:"status"`
	Attach          string `json:"attach"`
	CreatedAt       int64  `json:"createdAt"`
	CompletedAt     int64  `json:"completedAt"`
}

type QueryAllResponse struct {
	Total    string  `json:"total"`
	PageNo   string  `json:"pageNo"`
	PageSize string  `json:"pageSize"`
	List     []Order `json:"list"`
}
