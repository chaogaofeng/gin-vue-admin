// 自动生成模板APP
package upay
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 应用 结构体  APP
type APP struct {
    global.GVA_MODEL
    AppId  string `json:"appId" form:"appId" gorm:"uniqueIndex;column:app_id;comment:应用标识;" binding:"required"`  //应用标识 
    AppSecret  string `json:"appSecret" form:"appSecret" gorm:"column:app_secret;comment:应用密钥;" binding:"required"`  //应用密钥 
    AppName  string `json:"appName" form:"appName" gorm:"column:app_name;comment:应用名称;"`  //应用名称 
    UserID  *int `json:"userID" form:"userID" gorm:"column:user_id;comment:用户;" binding:"required"`  //用户 
    Status  string `json:"status" form:"status" gorm:"default:0;column:status;comment:应用状态;" binding:"required"`  //应用状态 
}


// TableName 应用 APP自定义表名 upay_app
func (APP) TableName() string {
    return "upay_app"
}
