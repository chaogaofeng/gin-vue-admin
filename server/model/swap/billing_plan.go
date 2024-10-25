// 自动生成模板BillingPlan
package swap

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
	"math"
)

// 账单计划 结构体  BillingPlan
type BillingPlan struct {
	global.GVA_MODEL
	Name         string   `json:"name" form:"name" gorm:"column:name;comment:名称;" binding:"required"`                                                         //名称
	Price        *float64 `json:"price" form:"price" gorm:"column:price;type:decimal(10,2);comment:价格;" binding:"required"`                                   //价格
	Duration     *int     `json:"duration" form:"duration" gorm:"column:duration;comment:数量;" binding:"required"`                                             //数量
	DurationUnit string   `json:"durationUnit" form:"durationUnit" gorm:"column:duration_unit;type:enum('D','W','M','Q','Y');comment:单位;" binding:"required"` //单位
	DailyPrice   *float64 `json:"dailyPrice" form:"dailyPrice" gorm:"column:daily_price;type:decimal(10,2);comment:日价格;"`                                     //日价格
	Discount     float64  `json:"discount" form:"discount" gorm:"-"`
}

// TableName 账单计划 BillingPlan自定义表名 billing_plan
func (BillingPlan) TableName() string {
	return "billing_plan"
}

// BeforeCreate 钩子在创建前自动计算 DailyRate 并设置默认名称
func (plan *BillingPlan) BeforeCreate(tx *gorm.DB) (err error) {
	dailyPrice := calculateDailyRate(*plan)
	plan.DailyPrice = &dailyPrice

	if plan.Name == "" {
		plan.Name = fmt.Sprintf("%vU/%v", plan.Price, plan.DurationUnit)
	}
	return nil
}
func (plan *BillingPlan) BeforeUpdate(tx *gorm.DB) (err error) {
	dailyPrice := calculateDailyRate(*plan)
	plan.DailyPrice = &dailyPrice

	if plan.Name == "" {
		plan.Name = fmt.Sprintf("%vU/%v", plan.Price, plan.DurationUnit)
	}
	return nil
}

func (plan *BillingPlan) AfterFind(tx *gorm.DB) (err error) {
	plan.Discount = 0
	if plan.DurationUnit == "D" && *plan.Duration == 1 {
		return
	}

	var dailyPlan BillingPlan
	if err := tx.Where("duration_unit = ? AND duration = ?", "D", 1).First(&dailyPlan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	dailyPrice := *dailyPlan.DailyPrice
	if plan.DurationUnit != "D" {
		// 使用总价格除以日价格计算折扣比例
		discount := (dailyPrice - *plan.DailyPrice) / dailyPrice * 100
		plan.Discount = math.Round(discount*100) / 100
	} else {
		plan.Discount = 0
	}
	return nil
}

// 计算日价的辅助函数
func calculateDailyRate(plan BillingPlan) float64 {
	days := 0
	switch plan.DurationUnit {
	case "D":
		days = *plan.Duration
	case "W":
		days = *plan.Duration * 7
	case "M":
		days = *plan.Duration * 30
	case "Q":
		days = *plan.Duration * 90
	case "Y":
		days = *plan.Duration * 365
	}

	return *plan.Price / float64(days)
}
