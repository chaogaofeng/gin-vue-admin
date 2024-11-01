package swap

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BillingPlanRouter struct{}

// InitBillingPlanRouter 初始化 账单计划 路由信息
func (s *BillingPlanRouter) InitBillingPlanRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	billingPlanRouter := Router.Group("billingPlan").Use(middleware.OperationRecord())
	//billingPlanRouterWithoutRecord := Router.Group("billingPlan")
	billingPlanRouterWithoutAuth := PublicRouter.Group("billingPlan")
	{
		billingPlanRouter.POST("createBillingPlan", billingPlanApi.CreateBillingPlan)             // 新建账单计划
		billingPlanRouter.DELETE("deleteBillingPlan", billingPlanApi.DeleteBillingPlan)           // 删除账单计划
		billingPlanRouter.DELETE("deleteBillingPlanByIds", billingPlanApi.DeleteBillingPlanByIds) // 批量删除账单计划
		billingPlanRouter.PUT("updateBillingPlan", billingPlanApi.UpdateBillingPlan)              // 更新账单计划
	}
	{
		billingPlanRouterWithoutAuth.GET("findBillingPlan", billingPlanApi.FindBillingPlan)       // 根据ID获取账单计划
		billingPlanRouterWithoutAuth.GET("getBillingPlanList", billingPlanApi.GetBillingPlanList) // 获取账单计划列表
	}
	{
		billingPlanRouterWithoutAuth.GET("getBillingPlanPublic", billingPlanApi.GetBillingPlanPublic) // 账单计划开放接口
	}
}
