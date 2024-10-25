package swap

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/swap"
    swapReq "github.com/flipped-aurora/gin-vue-admin/server/model/swap/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type BillingPlanApi struct {}



// CreateBillingPlan 创建账单计划
// @Tags BillingPlan
// @Summary 创建账单计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body swap.BillingPlan true "创建账单计划"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /billingPlan/createBillingPlan [post]
func (billingPlanApi *BillingPlanApi) CreateBillingPlan(c *gin.Context) {
	var billingPlan swap.BillingPlan
	err := c.ShouldBindJSON(&billingPlan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = billingPlanService.CreateBillingPlan(&billingPlan)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteBillingPlan 删除账单计划
// @Tags BillingPlan
// @Summary 删除账单计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body swap.BillingPlan true "删除账单计划"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /billingPlan/deleteBillingPlan [delete]
func (billingPlanApi *BillingPlanApi) DeleteBillingPlan(c *gin.Context) {
	ID := c.Query("ID")
	err := billingPlanService.DeleteBillingPlan(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBillingPlanByIds 批量删除账单计划
// @Tags BillingPlan
// @Summary 批量删除账单计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /billingPlan/deleteBillingPlanByIds [delete]
func (billingPlanApi *BillingPlanApi) DeleteBillingPlanByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := billingPlanService.DeleteBillingPlanByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBillingPlan 更新账单计划
// @Tags BillingPlan
// @Summary 更新账单计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body swap.BillingPlan true "更新账单计划"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /billingPlan/updateBillingPlan [put]
func (billingPlanApi *BillingPlanApi) UpdateBillingPlan(c *gin.Context) {
	var billingPlan swap.BillingPlan
	err := c.ShouldBindJSON(&billingPlan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = billingPlanService.UpdateBillingPlan(billingPlan)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBillingPlan 用id查询账单计划
// @Tags BillingPlan
// @Summary 用id查询账单计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query swap.BillingPlan true "用id查询账单计划"
// @Success 200 {object} response.Response{data=swap.BillingPlan,msg=string} "查询成功"
// @Router /billingPlan/findBillingPlan [get]
func (billingPlanApi *BillingPlanApi) FindBillingPlan(c *gin.Context) {
	ID := c.Query("ID")
	rebillingPlan, err := billingPlanService.GetBillingPlan(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rebillingPlan, c)
}

// GetBillingPlanList 分页获取账单计划列表
// @Tags BillingPlan
// @Summary 分页获取账单计划列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query swapReq.BillingPlanSearch true "分页获取账单计划列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /billingPlan/getBillingPlanList [get]
func (billingPlanApi *BillingPlanApi) GetBillingPlanList(c *gin.Context) {
	var pageInfo swapReq.BillingPlanSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := billingPlanService.GetBillingPlanInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetBillingPlanPublic 不需要鉴权的账单计划接口
// @Tags BillingPlan
// @Summary 不需要鉴权的账单计划接口
// @accept application/json
// @Produce application/json
// @Param data query swapReq.BillingPlanSearch true "分页获取账单计划列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /billingPlan/getBillingPlanPublic [get]
func (billingPlanApi *BillingPlanApi) GetBillingPlanPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    billingPlanService.GetBillingPlanPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的账单计划接口信息",
    }, "获取成功", c)
}
