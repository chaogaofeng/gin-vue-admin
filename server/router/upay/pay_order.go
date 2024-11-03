package upay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PayOrderRouter struct{}

// InitPayOrderRouter 初始化 支付订单 路由信息
func (s *PayOrderRouter) InitPayOrderRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	payOrderRouter := Router.Group("payOrder").Use(middleware.OperationRecord())
	payOrderRouterWithoutRecord := Router.Group("payOrder")
	payOrderRouterWithoutAuth := PublicRouter.Group("payOrder")
	{
		payOrderRouter.POST("createPayOrder", payOrderApi.CreatePayOrder)             // 新建支付订单
		payOrderRouter.DELETE("deletePayOrder", payOrderApi.DeletePayOrder)           // 删除支付订单
		payOrderRouter.DELETE("deletePayOrderByIds", payOrderApi.DeletePayOrderByIds) // 批量删除支付订单
		payOrderRouter.PUT("updatePayOrder", payOrderApi.UpdatePayOrder)              // 更新支付订单
	}
	{
		payOrderRouterWithoutRecord.GET("findPayOrder", payOrderApi.FindPayOrder)       // 根据ID获取支付订单
		payOrderRouterWithoutRecord.GET("getPayOrderList", payOrderApi.GetPayOrderList) // 获取支付订单列表
	}
	{
		payOrderRouterWithoutAuth.GET("getPayOrderDataSource", payOrderApi.GetPayOrderDataSource) // 获取支付订单数据源
		payOrderRouterWithoutAuth.GET("getPayOrderPublic", payOrderApi.GetPayOrderPublic)         // 支付订单开放接口
	}
}
