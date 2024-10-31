import service from '@/utils/request'
// @Tags PayOrder
// @Summary 创建支付订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayOrder true "创建支付订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /payOrder/createPayOrder [post]
export const createPayOrder = (data) => {
  return service({
    url: '/payOrder/createPayOrder',
    method: 'post',
    data
  })
}

// @Tags PayOrder
// @Summary 删除支付订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayOrder true "删除支付订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payOrder/deletePayOrder [delete]
export const deletePayOrder = (params) => {
  return service({
    url: '/payOrder/deletePayOrder',
    method: 'delete',
    params
  })
}

// @Tags PayOrder
// @Summary 批量删除支付订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除支付订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payOrder/deletePayOrder [delete]
export const deletePayOrderByIds = (params) => {
  return service({
    url: '/payOrder/deletePayOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags PayOrder
// @Summary 更新支付订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayOrder true "更新支付订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payOrder/updatePayOrder [put]
export const updatePayOrder = (data) => {
  return service({
    url: '/payOrder/updatePayOrder',
    method: 'put',
    data
  })
}

// @Tags PayOrder
// @Summary 用id查询支付订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PayOrder true "用id查询支付订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payOrder/findPayOrder [get]
export const findPayOrder = (params) => {
  return service({
    url: '/payOrder/findPayOrder',
    method: 'get',
    params
  })
}

// @Tags PayOrder
// @Summary 分页获取支付订单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取支付订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payOrder/getPayOrderList [get]
export const getPayOrderList = (params) => {
  return service({
    url: '/payOrder/getPayOrderList',
    method: 'get',
    params
  })
}
// @Tags PayOrder
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payOrder/findPayOrderDataSource [get]
export const getPayOrderDataSource = () => {
  return service({
    url: '/payOrder/getPayOrderDataSource',
    method: 'get',
  })
}

// @Tags PayOrder
// @Summary 不需要鉴权的支付订单接口
// @accept application/json
// @Produce application/json
// @Param data query upayReq.PayOrderSearch true "分页获取支付订单列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /payOrder/getPayOrderPublic [get]
export const getPayOrderPublic = () => {
  return service({
    url: '/payOrder/getPayOrderPublic',
    method: 'get',
  })
}
