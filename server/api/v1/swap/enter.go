package swap

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BillingPlanApi
	BillApi
}

var (
	billingPlanService = service.ServiceGroupApp.SwapServiceGroup.BillingPlanService
	billService        = service.ServiceGroupApp.SwapServiceGroup.BillService
)
