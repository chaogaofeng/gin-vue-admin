package upay

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	APPApi
	WalletAddressApi
	PayOrderApi
}

var (
	appService           = service.ServiceGroupApp.UpayServiceGroup.APPService
	walletAddressService = service.ServiceGroupApp.UpayServiceGroup.WalletAddressService
	payOrderService      = service.ServiceGroupApp.UpayServiceGroup.PayOrderService
)
