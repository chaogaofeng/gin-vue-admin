package swap

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	BillingPlanRouter
	BillRouter
}

var (
	billingPlanApi = api.ApiGroupApp.SwapApiGroup.BillingPlanApi
	billApi        = api.ApiGroupApp.SwapApiGroup.BillApi
)
