package upay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type APPRouter struct {}

// InitAPPRouter 初始化 应用 路由信息
func (s *APPRouter) InitAPPRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	appRouter := Router.Group("app").Use(middleware.OperationRecord())
	appRouterWithoutRecord := Router.Group("app")
	appRouterWithoutAuth := PublicRouter.Group("app")
	{
		appRouter.POST("createAPP", appApi.CreateAPP)   // 新建应用
		appRouter.DELETE("deleteAPP", appApi.DeleteAPP) // 删除应用
		appRouter.DELETE("deleteAPPByIds", appApi.DeleteAPPByIds) // 批量删除应用
		appRouter.PUT("updateAPP", appApi.UpdateAPP)    // 更新应用
	}
	{
		appRouterWithoutRecord.GET("findAPP", appApi.FindAPP)        // 根据ID获取应用
		appRouterWithoutRecord.GET("getAPPList", appApi.GetAPPList)  // 获取应用列表
	}
	{
	    appRouterWithoutAuth.GET("getAPPDataSource", appApi.GetAPPDataSource)  // 获取应用数据源
	    appRouterWithoutAuth.GET("getAPPPublic", appApi.GetAPPPublic)  // 应用开放接口
	}
}
