// 自动生成模板Bill
package swap

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 账单 结构体  Bill
type Bill struct {
	global.GVA_MODEL
	UserID     *int       `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户;"`                                     //账单计划
	No         string     `json:"no" form:"no" gorm:"column:no;comment:账单编号;"`                                                  //账单编号
	OrderNo    string     `json:"order_no" form:"order_no" gorm:"column:order_no;comment:订单编号;"`                                //订单编号
	PlanID     *int       `json:"plan_id" form:"plan_id" gorm:"column:plan_id;comment:账单计划;"`                                   //账单计划
	Count      *int       `json:"count" form:"count" gorm:"column:count;comment:账单计划数量;"`                                       //账单计划数量
	Address    string     `json:"address" form:"address" gorm:"column:address;comment:收款地址;"`                                   //交易哈希
	Hash       string     `json:"hash" form:"hash" gorm:"column:hash;comment:交易哈希;"`                                            //交易哈希
	ExpireAt   *time.Time `json:"expire_at" form:"expire_at" gorm:"column:expire_at;comment:开始时间;"`                             //开始时间
	StartAt    *time.Time `json:"start_at" form:"start_at" gorm:"column:start_at;comment:开始时间;"`                                //开始时间
	EndAt      *time.Time `json:"end_at" form:"end_at" gorm:"column:end_at;comment:结束时间;"`                                      //结束时间
	Status     *int       `json:"status" form:"status" gorm:"default:0;column:status;comment:账单状态; 0 待支付 1 支付成功 2 支付超时 3 支付失败"` //账单状态
	Amount     *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:账单金额;"`                                      //账单金额
	CompleteAt *time.Time `json:"complete_at" form:"complete_at" gorm:"column:complete_at;comment:完成时间;"`                       //完成时间
}

// TableName 账单 Bill自定义表名 bill
func (Bill) TableName() string {
	return "bill"
}
