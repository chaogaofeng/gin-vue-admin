package upay

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AmountGenerator struct {
	global.GVA_MODEL
	BaseValue    decimal.Decimal `gorm:"unique;column:base_value;type:decimal(10,4)"`
	CurrentValue decimal.Decimal `gorm:"column:current_value;type:decimal(10,4)"`
}

func (AmountGenerator) TableName() string {
	return "upay_amount_generator"
}

func getNextAmount(db *gorm.DB, baseValue, increment, upperLimit decimal.Decimal) (decimal.Decimal, error) {
	var amount AmountGenerator
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&amount, "base_value = ?", baseValue).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				amount = AmountGenerator{
					BaseValue:    baseValue,
					CurrentValue: baseValue,
				}
				return tx.Create(&amount).Error
			} else {
				return err
			}
		}

		nextDecimal := amount.CurrentValue.Add(increment)
		if nextDecimal.GreaterThan(upperLimit) {
			nextDecimal = baseValue
		}
		amount.CurrentValue = nextDecimal
		if err := tx.Save(&amount).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return decimal.Decimal{}, err
	}

	return amount.CurrentValue, nil
}
