package upay

import (
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AmountGenerator struct {
	ID             uint            `gorm:"primaryKey"`
	BaseValue      decimal.Decimal // 存储整数部分
	CurrentDecimal decimal.Decimal `gorm:"type:decimal(10,4)"` // 存储小数部分
}

func (AmountGenerator) TableName() string {
	return "upay_amount_generator"
}

func getNextAmount(db *gorm.DB, baseValue, upperLimit decimal.Decimal) (decimal.Decimal, error) {
	var amount AmountGenerator
	err := db.Transaction(func(tx *gorm.DB) error {
		// 获取当前金额
		if err := tx.First(&amount, "base_value = ?", baseValue).Error; err != nil {
			// 如果没有找到记录，返回初始金额
			if errors.Is(err, gorm.ErrRecordNotFound) {
				amount = AmountGenerator{
					BaseValue:      baseValue,
					CurrentDecimal: baseValue,
				}
				// 更新金额
				return tx.Create(&amount).Error
			} else {
				return err
			}
		}

		increment := decimal.NewFromFloat(0.0001) // 固定增量
		nextDecimal := amount.CurrentDecimal.Add(increment)
		if nextDecimal.GreaterThan(upperLimit) {
			nextDecimal = baseValue
		}
		// 更新金额
		amount.CurrentDecimal = nextDecimal
		if err := tx.Save(&amount).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return decimal.Decimal{}, err
	}

	return amount.CurrentDecimal, nil
}
