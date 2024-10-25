package swap

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/swap"
	swapReq "github.com/flipped-aurora/gin-vue-admin/server/model/swap/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type BillApi struct{}

// CreateBill 创建账单
// @Tags Bill
// @Summary 创建账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body swap.Bill true "创建账单"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /bill/createBill [post]
func (billApi *BillApi) CreateBill(c *gin.Context) {
	var bill swap.Bill
	err := c.ShouldBindJSON(&bill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user_id := int(utils.GetUserID(c))
	bill.UserID = &user_id
	billPlan, err := billingPlanService.GetBillingPlan(strconv.Itoa(*bill.PlanID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	amount := *billPlan.Price * float64(*bill.Count)
	// TODO
	bill.Amount = &amount
	expire_at := time.Now().Add(10 * time.Minute)
	bill.ExpireAt = &expire_at
	err = billService.CreateBill(&bill)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(bill, "创建成功", c)
}

// DeleteBill 删除账单
// @Tags Bill
// @Summary 删除账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body swap.Bill true "删除账单"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /bill/deleteBill [delete]
func (billApi *BillApi) DeleteBill(c *gin.Context) {
	ID := c.Query("ID")
	err := billService.DeleteBill(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBillByIds 批量删除账单
// @Tags Bill
// @Summary 批量删除账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /bill/deleteBillByIds [delete]
func (billApi *BillApi) DeleteBillByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := billService.DeleteBillByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBill 更新账单
// @Tags Bill
// @Summary 更新账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body swap.Bill true "更新账单"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bill/updateBill [put]
func (billApi *BillApi) UpdateBill(c *gin.Context) {
	var bill swap.Bill
	err := c.ShouldBindJSON(&bill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = billService.UpdateBill(bill)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBill 用id查询账单
// @Tags Bill
// @Summary 用id查询账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query swap.Bill true "用id查询账单"
// @Success 200 {object} response.Response{data=swap.Bill,msg=string} "查询成功"
// @Router /bill/findBill [get]
func (billApi *BillApi) FindBill(c *gin.Context) {
	ID := c.Query("ID")
	rebill, err := billService.GetBill(ID)
	if no := c.Query("no"); len(no) > 0 {
		rebill, err = billService.GetBillByNo(no)
	}
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rebill, c)
}

// GetBillList 分页获取账单列表
// @Tags Bill
// @Summary 分页获取账单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query swapReq.BillSearch true "分页获取账单列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /bill/getBillList [get]
func (billApi *BillApi) GetBillList(c *gin.Context) {
	var pageInfo swapReq.BillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := billService.GetBillInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetBillDataSource 获取Bill的数据源
// @Tags Bill
// @Summary 获取Bill的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /bill/getBillDataSource [get]
func (billApi *BillApi) GetBillDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	dataSource, err := billService.GetBillDataSource()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetBillPublic 不需要鉴权的账单接口
// @Tags Bill
// @Summary 不需要鉴权的账单接口
// @accept application/json
// @Produce application/json
// @Param data query swapReq.BillSearch true "分页获取账单列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bill/getBillPublic [get]
func (billApi *BillApi) GetBillPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	billService.GetBillPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的账单接口信息",
	}, "获取成功", c)
}
