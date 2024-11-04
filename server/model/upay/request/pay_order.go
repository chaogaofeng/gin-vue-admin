package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PayOrderSearch struct {
	StartCreatedAt   *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt     *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Status           string     `json:"status" form:"status" `
	Amount           *float64   `json:"crypto" form:"crypto" `
	ActualAmount     *float64   `json:"actualCrypto" form:"actualCrypto" `
	ChainType        string     `json:"chainType" form:"chainType" `
	AppID            string     `json:"appId" form:"appId" `
	Receiver         string     `json:"receiver" form:"receiver" `
	Hash             string     `json:"hash" form:"hash" `
	OrderNo          string     `json:"orderNo" form:"orderNo" `
	MerchantOrderNo  string     `json:"merchantOrderNo" form:"merchantOrderNo" `
	StartCompletedAt *time.Time `json:"startCompletedAt" form:"startCompletedAt"`
	EndCompletedAt   *time.Time `json:"endCompletedAt" form:"endCompletedAt"`
	RiskLevel        string     `json:"riskLevel" form:"riskLevel" `
	UserID           uint       `json:"userID" form:"userID" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
