import service from '@/utils/request'
// @Tags Bill
// @Summary 创建账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Bill true "创建账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bill/createBill [post]
export const createBill = (data) => {
  return service({
    url: '/bill/createBill',
    method: 'post',
    data
  })
}

// @Tags Bill
// @Summary 删除账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Bill true "删除账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bill/deleteBill [delete]
export const deleteBill = (params) => {
  return service({
    url: '/bill/deleteBill',
    method: 'delete',
    params
  })
}

// @Tags Bill
// @Summary 批量删除账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bill/deleteBill [delete]
export const deleteBillByIds = (params) => {
  return service({
    url: '/bill/deleteBillByIds',
    method: 'delete',
    params
  })
}

// @Tags Bill
// @Summary 更新账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Bill true "更新账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bill/updateBill [put]
export const updateBill = (data) => {
  return service({
    url: '/bill/updateBill',
    method: 'put',
    data
  })
}

// @Tags Bill
// @Summary 用id查询账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Bill true "用id查询账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bill/findBill [get]
export const findBill = (params) => {
  return service({
    url: '/bill/findBill',
    method: 'get',
    params
  })
}

// @Tags Bill
// @Summary 分页获取账单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取账单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bill/getBillList [get]
export const getBillList = (params) => {
  return service({
    url: '/bill/getBillList',
    method: 'get',
    params
  })
}
// @Tags Bill
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bill/findBillDataSource [get]
export const getBillDataSource = () => {
  return service({
    url: '/bill/getBillDataSource',
    method: 'get',
  })
}

// @Tags Bill
// @Summary 不需要鉴权的账单接口
// @accept application/json
// @Produce application/json
// @Param data query swapReq.BillSearch true "分页获取账单列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bill/getBillPublic [get]
export const getBillPublic = () => {
  return service({
    url: '/bill/getBillPublic',
    method: 'get',
  })
}
