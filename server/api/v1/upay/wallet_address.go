package upay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/upay"
	upayReq "github.com/flipped-aurora/gin-vue-admin/server/model/upay/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WalletAddressApi struct{}

// CreateWalletAddress 创建收款钱包
// @Tags WalletAddress
// @Summary 创建收款钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body upay.WalletAddress true "创建收款钱包"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /walletAddress/createWalletAddress [post]
func (walletAddressApi *WalletAddressApi) CreateWalletAddress(c *gin.Context) {
	var walletAddress upay.WalletAddress
	err := c.ShouldBindJSON(&walletAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	walletAddress.UserID = utils.GetUserID(c)
	err = walletAddressService.CreateWalletAddress(&walletAddress)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.createSuccess"), c)
}

// DeleteWalletAddress 删除收款钱包
// @Tags WalletAddress
// @Summary 删除收款钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body upay.WalletAddress true "删除收款钱包"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /walletAddress/deleteWalletAddress [delete]
func (walletAddressApi *WalletAddressApi) DeleteWalletAddress(c *gin.Context) {
	ID := c.Query("ID")
	err := walletAddressService.DeleteWalletAddress(ID)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// DeleteWalletAddressByIds 批量删除收款钱包
// @Tags WalletAddress
// @Summary 批量删除收款钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /walletAddress/deleteWalletAddressByIds [delete]
func (walletAddressApi *WalletAddressApi) DeleteWalletAddressByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := walletAddressService.DeleteWalletAddressByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("system.sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("system.sys_operation_record.batchDeleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("system.sys_operation_record.batchDeleteSuccess"), c)
}

// UpdateWalletAddress 更新收款钱包
// @Tags WalletAddress
// @Summary 更新收款钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body upay.WalletAddress true "更新收款钱包"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /walletAddress/updateWalletAddress [put]
func (walletAddressApi *WalletAddressApi) UpdateWalletAddress(c *gin.Context) {
	var walletAddress upay.WalletAddress
	err := c.ShouldBindJSON(&walletAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = walletAddressService.UpdateWalletAddress(walletAddress)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.updateSuccess"), c)
}

// FindWalletAddress 用id查询收款钱包
// @Tags WalletAddress
// @Summary 用id查询收款钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query upay.WalletAddress true "用id查询收款钱包"
// @Success 200 {object} response.Response{data=upay.WalletAddress,msg=string} "查询成功"
// @Router /walletAddress/findWalletAddress [get]
func (walletAddressApi *WalletAddressApi) FindWalletAddress(c *gin.Context) {
	ID := c.Query("ID")
	rewalletAddress, err := walletAddressService.GetWalletAddress(ID)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithData(rewalletAddress, c)
}

// GetWalletAddressList 分页获取收款钱包列表
// @Tags WalletAddress
// @Summary 分页获取收款钱包列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query upayReq.WalletAddressSearch true "分页获取收款钱包列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /walletAddress/getWalletAddressList [get]
func (walletAddressApi *WalletAddressApi) GetWalletAddressList(c *gin.Context) {
	var pageInfo upayReq.WalletAddressSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := walletAddressService.GetWalletAddressInfoList(pageInfo)
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

// GetWalletAddressDataSource 获取WalletAddress的数据源
// @Tags WalletAddress
// @Summary 获取WalletAddress的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /walletAddress/getWalletAddressDataSource [get]
func (walletAddressApi *WalletAddressApi) GetWalletAddressDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	dataSource, err := walletAddressService.GetWalletAddressDataSource()
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetWalletAddressPublic 不需要鉴权的收款钱包接口
// @Tags WalletAddress
// @Summary 不需要鉴权的收款钱包接口
// @accept application/json
// @Produce application/json
// @Param data query upayReq.WalletAddressSearch true "分页获取收款钱包列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /walletAddress/getWalletAddressPublic [get]
func (walletAddressApi *WalletAddressApi) GetWalletAddressPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	walletAddressService.GetWalletAddressPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的收款钱包接口信息",
	}, "获取成功", c)
}
