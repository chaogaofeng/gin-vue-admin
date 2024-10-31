import service from '@/utils/request'
// @Tags WalletAddress
// @Summary 创建收款钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WalletAddress true "创建收款钱包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /walletAddress/createWalletAddress [post]
export const createWalletAddress = (data) => {
  return service({
    url: '/walletAddress/createWalletAddress',
    method: 'post',
    data
  })
}

// @Tags WalletAddress
// @Summary 删除收款钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WalletAddress true "删除收款钱包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /walletAddress/deleteWalletAddress [delete]
export const deleteWalletAddress = (params) => {
  return service({
    url: '/walletAddress/deleteWalletAddress',
    method: 'delete',
    params
  })
}

// @Tags WalletAddress
// @Summary 批量删除收款钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除收款钱包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /walletAddress/deleteWalletAddress [delete]
export const deleteWalletAddressByIds = (params) => {
  return service({
    url: '/walletAddress/deleteWalletAddressByIds',
    method: 'delete',
    params
  })
}

// @Tags WalletAddress
// @Summary 更新收款钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WalletAddress true "更新收款钱包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /walletAddress/updateWalletAddress [put]
export const updateWalletAddress = (data) => {
  return service({
    url: '/walletAddress/updateWalletAddress',
    method: 'put',
    data
  })
}

// @Tags WalletAddress
// @Summary 用id查询收款钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WalletAddress true "用id查询收款钱包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /walletAddress/findWalletAddress [get]
export const findWalletAddress = (params) => {
  return service({
    url: '/walletAddress/findWalletAddress',
    method: 'get',
    params
  })
}

// @Tags WalletAddress
// @Summary 分页获取收款钱包列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取收款钱包列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /walletAddress/getWalletAddressList [get]
export const getWalletAddressList = (params) => {
  return service({
    url: '/walletAddress/getWalletAddressList',
    method: 'get',
    params
  })
}
// @Tags WalletAddress
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /walletAddress/findWalletAddressDataSource [get]
export const getWalletAddressDataSource = () => {
  return service({
    url: '/walletAddress/getWalletAddressDataSource',
    method: 'get',
  })
}

// @Tags WalletAddress
// @Summary 不需要鉴权的收款钱包接口
// @accept application/json
// @Produce application/json
// @Param data query upayReq.WalletAddressSearch true "分页获取收款钱包列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /walletAddress/getWalletAddressPublic [get]
export const getWalletAddressPublic = () => {
  return service({
    url: '/walletAddress/getWalletAddressPublic',
    method: 'get',
  })
}
