package swap

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BillRouter struct {}

// InitBillRouter 初始化 账单 路由信息
func (s *BillRouter) InitBillRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	billRouter := Router.Group("bill").Use(middleware.OperationRecord())
	billRouterWithoutRecord := Router.Group("bill")
	billRouterWithoutAuth := PublicRouter.Group("bill")
	{
		billRouter.POST("createBill", billApi.CreateBill)   // 新建账单
		billRouter.DELETE("deleteBill", billApi.DeleteBill) // 删除账单
		billRouter.DELETE("deleteBillByIds", billApi.DeleteBillByIds) // 批量删除账单
		billRouter.PUT("updateBill", billApi.UpdateBill)    // 更新账单
	}
	{
		billRouterWithoutRecord.GET("findBill", billApi.FindBill)        // 根据ID获取账单
		billRouterWithoutRecord.GET("getBillList", billApi.GetBillList)  // 获取账单列表
	}
	{
	    billRouterWithoutAuth.GET("getBillDataSource", billApi.GetBillDataSource)  // 获取账单数据源
	    billRouterWithoutAuth.GET("getBillPublic", billApi.GetBillPublic)  // 账单开放接口
	}
}
