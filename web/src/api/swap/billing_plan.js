import service from '@/utils/request'
// @Tags BillingPlan
// @Summary 创建账单计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BillingPlan true "创建账单计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /billingPlan/createBillingPlan [post]
export const createBillingPlan = (data) => {
  return service({
    url: '/billingPlan/createBillingPlan',
    method: 'post',
    data
  })
}

// @Tags BillingPlan
// @Summary 删除账单计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BillingPlan true "删除账单计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /billingPlan/deleteBillingPlan [delete]
export const deleteBillingPlan = (params) => {
  return service({
    url: '/billingPlan/deleteBillingPlan',
    method: 'delete',
    params
  })
}

// @Tags BillingPlan
// @Summary 批量删除账单计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除账单计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /billingPlan/deleteBillingPlan [delete]
export const deleteBillingPlanByIds = (params) => {
  return service({
    url: '/billingPlan/deleteBillingPlanByIds',
    method: 'delete',
    params
  })
}

// @Tags BillingPlan
// @Summary 更新账单计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BillingPlan true "更新账单计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /billingPlan/updateBillingPlan [put]
export const updateBillingPlan = (data) => {
  return service({
    url: '/billingPlan/updateBillingPlan',
    method: 'put',
    data
  })
}

// @Tags BillingPlan
// @Summary 用id查询账单计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BillingPlan true "用id查询账单计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /billingPlan/findBillingPlan [get]
export const findBillingPlan = (params) => {
  return service({
    url: '/billingPlan/findBillingPlan',
    method: 'get',
    params
  })
}

// @Tags BillingPlan
// @Summary 分页获取账单计划列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取账单计划列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /billingPlan/getBillingPlanList [get]
export const getBillingPlanList = (params) => {
  return service({
    url: '/billingPlan/getBillingPlanList',
    method: 'get',
    params
  })
}

// @Tags BillingPlan
// @Summary 不需要鉴权的账单计划接口
// @accept application/json
// @Produce application/json
// @Param data query swapReq.BillingPlanSearch true "分页获取账单计划列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /billingPlan/getBillingPlanPublic [get]
export const getBillingPlanPublic = () => {
  return service({
    url: '/billingPlan/getBillingPlanPublic',
    method: 'get',
  })
}
