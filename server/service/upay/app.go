package upay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/upay"
	upayReq "github.com/flipped-aurora/gin-vue-admin/server/model/upay/request"
)

type APPService struct{}

// CreateAPP 创建应用记录
// Author [yourname](https://github.com/yourname)
func (appService *APPService) CreateAPP(app *upay.APP) (err error) {
	err = global.GVA_DB.Create(app).Error
	return err
}

// DeleteAPP 删除应用记录
// Author [yourname](https://github.com/yourname)
func (appService *APPService) DeleteAPP(ID string) (err error) {
	err = global.GVA_DB.Delete(&upay.APP{}, "id = ?", ID).Error
	return err
}

// DeleteAPPByIds 批量删除应用记录
// Author [yourname](https://github.com/yourname)
func (appService *APPService) DeleteAPPByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]upay.APP{}, "id in ?", IDs).Error
	return err
}

// UpdateAPP 更新应用记录
// Author [yourname](https://github.com/yourname)
func (appService *APPService) UpdateAPP(app upay.APP) (err error) {
	err = global.GVA_DB.Model(&upay.APP{}).Where("id = ?", app.ID).Updates(&app).Error
	return err
}

// GetAPP 根据ID获取应用记录
// Author [yourname](https://github.com/yourname)
func (appService *APPService) GetAPP(ID string) (app upay.APP, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&app).Error
	return
}

// GetAPPInfoList 分页获取应用记录
// Author [yourname](https://github.com/yourname)
func (appService *APPService) GetAPPInfoList(info upayReq.APPSearch) (list []upay.APP, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&upay.APP{})
	var apps []upay.APP
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AppId != "" {
		db = db.Where("app_id = ?", info.AppId)
	}
	if info.AppSecret != "" {
		db = db.Where("app_secret = ?", info.AppSecret)
	}
	if info.AppName != "" {
		db = db.Where("app_name LIKE ?", "%"+info.AppName+"%")
	}
	if info.UserID != nil {
		db = db.Where("user_id = ?", info.UserID)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&apps).Error
	return apps, total, err
}
func (appService *APPService) GetAPPDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	userID := make([]map[string]any, 0)

	global.GVA_DB.Table("sys_users").Select("username as label,id as value").Scan(&userID)
	res["userID"] = userID
	return
}
func (appService *APPService) GetAPPPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
