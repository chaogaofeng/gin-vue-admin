package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type WalletAddressSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Address        string     `json:"address" form:"address" `
	Name           string     `json:"name" form:"name" `
	ChainType      string     `json:"chainType" form:"chainType" `
	Status         string     `json:"status" form:"status" `
	UserID         uint       `json:"userID" form:"userID" `
	request.PageInfo
}
