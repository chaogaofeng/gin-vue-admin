package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type APPSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	AppId          string     `json:"appId" form:"appId" `
	AppSecret      string     `json:"appSecret" form:"appSecret" `
	AppName        string     `json:"appName" form:"appName" `
	UserID         *int       `json:"userID" form:"userID" `
	Status         string     `json:"status" form:"status" `
	request.PageInfo
}
