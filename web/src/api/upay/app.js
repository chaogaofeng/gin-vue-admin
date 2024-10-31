import service from '@/utils/request'
// @Tags APP
// @Summary 创建应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.APP true "创建应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /app/createAPP [post]
export const createAPP = (data) => {
  return service({
    url: '/app/createAPP',
    method: 'post',
    data
  })
}

// @Tags APP
// @Summary 删除应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.APP true "删除应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/deleteAPP [delete]
export const deleteAPP = (params) => {
  return service({
    url: '/app/deleteAPP',
    method: 'delete',
    params
  })
}

// @Tags APP
// @Summary 批量删除应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/deleteAPP [delete]
export const deleteAPPByIds = (params) => {
  return service({
    url: '/app/deleteAPPByIds',
    method: 'delete',
    params
  })
}

// @Tags APP
// @Summary 更新应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.APP true "更新应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/updateAPP [put]
export const updateAPP = (data) => {
  return service({
    url: '/app/updateAPP',
    method: 'put',
    data
  })
}

// @Tags APP
// @Summary 用id查询应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.APP true "用id查询应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/findAPP [get]
export const findAPP = (params) => {
  return service({
    url: '/app/findAPP',
    method: 'get',
    params
  })
}

// @Tags APP
// @Summary 分页获取应用列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取应用列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/getAPPList [get]
export const getAPPList = (params) => {
  return service({
    url: '/app/getAPPList',
    method: 'get',
    params
  })
}
// @Tags APP
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/findAPPDataSource [get]
export const getAPPDataSource = () => {
  return service({
    url: '/app/getAPPDataSource',
    method: 'get',
  })
}

// @Tags APP
// @Summary 不需要鉴权的应用接口
// @accept application/json
// @Produce application/json
// @Param data query upayReq.APPSearch true "分页获取应用列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /app/getAPPPublic [get]
export const getAPPPublic = () => {
  return service({
    url: '/app/getAPPPublic',
    method: 'get',
  })
}
