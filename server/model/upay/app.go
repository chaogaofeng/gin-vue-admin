// 自动生成模板APP
package upay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

// 应用 结构体  APP
type APP struct {
	global.GVA_MODEL
	AppName   string `json:"appName" form:"appName" gorm:"column:app_name;comment:应用名称;"`              //应用名称
	AppId     string `json:"appId" form:"appId" gorm:"uniqueIndex;column:app_id;comment:应用标识;"`        //应用标识
	AppSecret string `json:"appSecret" form:"appSecret" gorm:"unique;column:app_secret;comment:应用密钥;"` //应用密钥
	Status    string `json:"status" form:"status" gorm:"default:0;column:status;comment:应用状态;"`        //应用状态
	UserID    uint   `json:"userID" form:"userID" gorm:"column:user_id;comment:用户;"`                   //用户
}

// TableName 应用 APP自定义表名 upay_app
func (APP) TableName() string {
	return "upay_app"
}

func (app *APP) BeforeCreate(db *gorm.DB) error {
	if len(app.AppId) == 0 {
		for {
			appId, err := utils.GenerateRandomString(8)
			if err != nil {
				return err
			}
			var count int64
			if err := db.Model(&APP{}).Where("app_id = ?", appId).Count(&count).Error; err != nil {
				return err
			}
			if count == 0 {
				app.AppId = appId
				break
			}
		}
	}
	if len(app.AppSecret) == 0 {
		for {
			appSecret, err := utils.GenerateRandomString(16)
			if err != nil {
				return err
			}
			var count int64
			if err := db.Model(&APP{}).Where("app_secret = ?", appSecret).Count(&count).Error; err != nil {
				return err
			}
			if count == 0 {
				app.AppSecret = appSecret
				break
			}
		}
	}
	//if len(app.AppName) == 0 {
	//	name, err := utils.GenerateRandomString(6)
	//	if err != nil {
	//		return err
	//	}
	//	app.AppName = name
	//
	//}
	return nil
}
