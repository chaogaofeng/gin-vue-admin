package swap

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/swap"
    swapReq "github.com/flipped-aurora/gin-vue-admin/server/model/swap/request"
)

type BillingPlanService struct {}
// CreateBillingPlan 创建账单计划记录
// Author [yourname](https://github.com/yourname)
func (billingPlanService *BillingPlanService) CreateBillingPlan(billingPlan *swap.BillingPlan) (err error) {
	err = global.GVA_DB.Create(billingPlan).Error
	return err
}

// DeleteBillingPlan 删除账单计划记录
// Author [yourname](https://github.com/yourname)
func (billingPlanService *BillingPlanService)DeleteBillingPlan(ID string) (err error) {
	err = global.GVA_DB.Delete(&swap.BillingPlan{},"id = ?",ID).Error
	return err
}

// DeleteBillingPlanByIds 批量删除账单计划记录
// Author [yourname](https://github.com/yourname)
func (billingPlanService *BillingPlanService)DeleteBillingPlanByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]swap.BillingPlan{},"id in ?",IDs).Error
	return err
}

// UpdateBillingPlan 更新账单计划记录
// Author [yourname](https://github.com/yourname)
func (billingPlanService *BillingPlanService)UpdateBillingPlan(billingPlan swap.BillingPlan) (err error) {
	err = global.GVA_DB.Model(&swap.BillingPlan{}).Where("id = ?",billingPlan.ID).Updates(&billingPlan).Error
	return err
}

// GetBillingPlan 根据ID获取账单计划记录
// Author [yourname](https://github.com/yourname)
func (billingPlanService *BillingPlanService)GetBillingPlan(ID string) (billingPlan swap.BillingPlan, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&billingPlan).Error
	return
}

// GetBillingPlanInfoList 分页获取账单计划记录
// Author [yourname](https://github.com/yourname)
func (billingPlanService *BillingPlanService)GetBillingPlanInfoList(info swapReq.BillingPlanSearch) (list []swap.BillingPlan, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&swap.BillingPlan{})
    var billingPlans []swap.BillingPlan
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.DurationUnit != "" {
        db = db.Where("duration_unit = ?",info.DurationUnit)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["price"] = true
         	orderMap["duration"] = true
         	orderMap["duration_unit"] = true
         	orderMap["daily_price"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&billingPlans).Error
	return  billingPlans, total, err
}
func (billingPlanService *BillingPlanService)GetBillingPlanPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
