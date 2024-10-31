package upay

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/upay"
    upayReq "github.com/flipped-aurora/gin-vue-admin/server/model/upay/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type APPApi struct {}



// CreateAPP 创建应用
// @Tags APP
// @Summary 创建应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body upay.APP true "创建应用"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /app/createAPP [post]
func (appApi *APPApi) CreateAPP(c *gin.Context) {
	var app upay.APP
	err := c.ShouldBindJSON(&app)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = appService.CreateAPP(&app)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
		return
	}
    response.OkWithMessage(global.Translate("general.createSuccess"), c)
}

// DeleteAPP 删除应用
// @Tags APP
// @Summary 删除应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body upay.APP true "删除应用"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /app/deleteAPP [delete]
func (appApi *APPApi) DeleteAPP(c *gin.Context) {
	ID := c.Query("ID")
	err := appService.DeleteAPP(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// DeleteAPPByIds 批量删除应用
// @Tags APP
// @Summary 批量删除应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /app/deleteAPPByIds [delete]
func (appApi *APPApi) DeleteAPPByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := appService.DeleteAPPByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("system.sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("system.sys_operation_record.batchDeleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("system.sys_operation_record.batchDeleteSuccess"), c)
}

// UpdateAPP 更新应用
// @Tags APP
// @Summary 更新应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body upay.APP true "更新应用"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /app/updateAPP [put]
func (appApi *APPApi) UpdateAPP(c *gin.Context) {
	var app upay.APP
	err := c.ShouldBindJSON(&app)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = appService.UpdateAPP(app)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.updateSuccess"), c)
}

// FindAPP 用id查询应用
// @Tags APP
// @Summary 用id查询应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query upay.APP true "用id查询应用"
// @Success 200 {object} response.Response{data=upay.APP,msg=string} "查询成功"
// @Router /app/findAPP [get]
func (appApi *APPApi) FindAPP(c *gin.Context) {
	ID := c.Query("ID")
	reapp, err := appService.GetAPP(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithData(reapp, c)
}

// GetAPPList 分页获取应用列表
// @Tags APP
// @Summary 分页获取应用列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query upayReq.APPSearch true "分页获取应用列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /app/getAPPList [get]
func (appApi *APPApi) GetAPPList(c *gin.Context) {
	var pageInfo upayReq.APPSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := appService.GetAPPInfoList(pageInfo)
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
// GetAPPDataSource 获取APP的数据源
// @Tags APP
// @Summary 获取APP的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /app/getAPPDataSource [get]
func (appApi *APPApi) GetAPPDataSource(c *gin.Context) {
    // 此接口为获取数据源定义的数据
    dataSource, err := appService.GetAPPDataSource()
    if err != nil {
        global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
   		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
   		return
    }
   response.OkWithData(dataSource, c)
}

// GetAPPPublic 不需要鉴权的应用接口
// @Tags APP
// @Summary 不需要鉴权的应用接口
// @accept application/json
// @Produce application/json
// @Param data query upayReq.APPSearch true "分页获取应用列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /app/getAPPPublic [get]
func (appApi *APPApi) GetAPPPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    appService.GetAPPPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的应用接口信息",
    }, "获取成功", c)
}
