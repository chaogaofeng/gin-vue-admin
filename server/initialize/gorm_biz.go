package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/swap"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(swap.BillingPlan{}, swap.Bill{})
	if err != nil {
		return err
	}
	return nil
}
