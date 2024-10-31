package upay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WalletAddressRouter struct {}

// InitWalletAddressRouter 初始化 收款钱包 路由信息
func (s *WalletAddressRouter) InitWalletAddressRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	walletAddressRouter := Router.Group("walletAddress").Use(middleware.OperationRecord())
	walletAddressRouterWithoutRecord := Router.Group("walletAddress")
	walletAddressRouterWithoutAuth := PublicRouter.Group("walletAddress")
	{
		walletAddressRouter.POST("createWalletAddress", walletAddressApi.CreateWalletAddress)   // 新建收款钱包
		walletAddressRouter.DELETE("deleteWalletAddress", walletAddressApi.DeleteWalletAddress) // 删除收款钱包
		walletAddressRouter.DELETE("deleteWalletAddressByIds", walletAddressApi.DeleteWalletAddressByIds) // 批量删除收款钱包
		walletAddressRouter.PUT("updateWalletAddress", walletAddressApi.UpdateWalletAddress)    // 更新收款钱包
	}
	{
		walletAddressRouterWithoutRecord.GET("findWalletAddress", walletAddressApi.FindWalletAddress)        // 根据ID获取收款钱包
		walletAddressRouterWithoutRecord.GET("getWalletAddressList", walletAddressApi.GetWalletAddressList)  // 获取收款钱包列表
	}
	{
	    walletAddressRouterWithoutAuth.GET("getWalletAddressDataSource", walletAddressApi.GetWalletAddressDataSource)  // 获取收款钱包数据源
	    walletAddressRouterWithoutAuth.GET("getWalletAddressPublic", walletAddressApi.GetWalletAddressPublic)  // 收款钱包开放接口
	}
}
