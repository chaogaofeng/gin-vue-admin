package upay

import (
	"github.com/gin-gonic/gin"
)

type OpenRouter struct{}

// InitAPPRouter 初始化 应用 路由信息
func (s *OpenRouter) InitAPPRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	orderRouterWithoutAuth := PublicRouter.Group("open").Group("order")
	///v1/api/open/order/apply
	///v1/api/open/order/search
	///v1/api/open/order/queryall
	{
		orderRouterWithoutAuth.GET("apply", openApi.OrderApply)   // 获取应用数据源
		orderRouterWithoutAuth.GET("search", openApi.OrderSearch) // 应用开放接口
		orderRouterWithoutAuth.GET("queryall", openApi.QueryAll)  // 应用开放接口
	}
}
