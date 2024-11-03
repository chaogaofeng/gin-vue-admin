// 自动生成模板WalletAddress
package upay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 收款钱包 结构体  WalletAddress
type WalletAddress struct {
	global.GVA_MODEL
	Address   string `json:"address" form:"address" gorm:"uniqueIndex;column:address;comment:钱包地址;" binding:"required"`       //钱包地址
	Name      string `json:"name" form:"name" gorm:"column:name;comment:钱包名称;"`                                               //钱包名称
	ChainType string `json:"chainType" form:"chainType" gorm:"default:0;column:chain_type;comment:区块链类型;" binding:"required"` //区块链类型
	Status    string `json:"status" form:"status" gorm:"default:0;column:status;comment:地址状态;" binding:"required"`            //地址状态
	UserID    *int   `json:"userID" form:"userID" gorm:"column:user_id;comment:用户;"`                                          //用户
}

// TableName 收款钱包 WalletAddress自定义表名 upay_wallet_address
func (WalletAddress) TableName() string {
	return "upay_wallet_address"
}
