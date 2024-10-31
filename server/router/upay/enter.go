package upay

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	APPRouter
	WalletAddressRouter
	PayOrderRouter
}

var (
	appApi           = api.ApiGroupApp.UpayApiGroup.APPApi
	walletAddressApi = api.ApiGroupApp.UpayApiGroup.WalletAddressApi
	payOrderApi      = api.ApiGroupApp.UpayApiGroup.PayOrderApi
)
