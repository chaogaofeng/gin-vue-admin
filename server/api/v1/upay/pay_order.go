package upay

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/upay"
    upayReq "github.com/flipped-aurora/gin-vue-admin/server/model/upay/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PayOrderApi struct {}



// CreatePayOrder 创建支付订单
// @Tags PayOrder
// @Summary 创建支付订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body upay.PayOrder true "创建支付订单"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /payOrder/createPayOrder [post]
func (payOrderApi *PayOrderApi) CreatePayOrder(c *gin.Context) {
	var payOrder upay.PayOrder
	err := c.ShouldBindJSON(&payOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = payOrderService.CreatePayOrder(&payOrder)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
		return
	}
    response.OkWithMessage(global.Translate("general.createSuccess"), c)
}

// DeletePayOrder 删除支付订单
// @Tags PayOrder
// @Summary 删除支付订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body upay.PayOrder true "删除支付订单"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /payOrder/deletePayOrder [delete]
func (payOrderApi *PayOrderApi) DeletePayOrder(c *gin.Context) {
	ID := c.Query("ID")
	err := payOrderService.DeletePayOrder(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// DeletePayOrderByIds 批量删除支付订单
// @Tags PayOrder
// @Summary 批量删除支付订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /payOrder/deletePayOrderByIds [delete]
func (payOrderApi *PayOrderApi) DeletePayOrderByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := payOrderService.DeletePayOrderByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("system.sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("system.sys_operation_record.batchDeleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("system.sys_operation_record.batchDeleteSuccess"), c)
}

// UpdatePayOrder 更新支付订单
// @Tags PayOrder
// @Summary 更新支付订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body upay.PayOrder true "更新支付订单"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /payOrder/updatePayOrder [put]
func (payOrderApi *PayOrderApi) UpdatePayOrder(c *gin.Context) {
	var payOrder upay.PayOrder
	err := c.ShouldBindJSON(&payOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = payOrderService.UpdatePayOrder(payOrder)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.updateSuccess"), c)
}

// FindPayOrder 用id查询支付订单
// @Tags PayOrder
// @Summary 用id查询支付订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query upay.PayOrder true "用id查询支付订单"
// @Success 200 {object} response.Response{data=upay.PayOrder,msg=string} "查询成功"
// @Router /payOrder/findPayOrder [get]
func (payOrderApi *PayOrderApi) FindPayOrder(c *gin.Context) {
	ID := c.Query("ID")
	repayOrder, err := payOrderService.GetPayOrder(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithData(repayOrder, c)
}

// GetPayOrderList 分页获取支付订单列表
// @Tags PayOrder
// @Summary 分页获取支付订单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query upayReq.PayOrderSearch true "分页获取支付订单列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /payOrder/getPayOrderList [get]
func (payOrderApi *PayOrderApi) GetPayOrderList(c *gin.Context) {
	var pageInfo upayReq.PayOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := payOrderService.GetPayOrderInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
        response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, global.Translate("general.getDataSuccess"), c)
}
// GetPayOrderDataSource 获取PayOrder的数据源
// @Tags PayOrder
// @Summary 获取PayOrder的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /payOrder/getPayOrderDataSource [get]
func (payOrderApi *PayOrderApi) GetPayOrderDataSource(c *gin.Context) {
    // 此接口为获取数据源定义的数据
    dataSource, err := payOrderService.GetPayOrderDataSource()
    if err != nil {
        global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
   		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
   		return
    }
   response.OkWithData(dataSource, c)
}

// GetPayOrderPublic 不需要鉴权的支付订单接口
// @Tags PayOrder
// @Summary 不需要鉴权的支付订单接口
// @accept application/json
// @Produce application/json
// @Param data query upayReq.PayOrderSearch true "分页获取支付订单列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /payOrder/getPayOrderPublic [get]
func (payOrderApi *PayOrderApi) GetPayOrderPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    payOrderService.GetPayOrderPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的支付订单接口信息",
    }, "获取成功", c)
}
