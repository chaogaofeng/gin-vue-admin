package upay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/upay"
	upayReq "github.com/flipped-aurora/gin-vue-admin/server/model/upay/request"
)

type PayOrderService struct{}

// CreatePayOrder 创建支付订单记录
// Author [yourname](https://github.com/yourname)
func (payOrderService *PayOrderService) CreatePayOrder(payOrder *upay.PayOrder) (err error) {
	err = global.GVA_DB.Create(payOrder).Error
	return err
}

// DeletePayOrder 删除支付订单记录
// Author [yourname](https://github.com/yourname)
func (payOrderService *PayOrderService) DeletePayOrder(ID string) (err error) {
	err = global.GVA_DB.Delete(&upay.PayOrder{}, "id = ?", ID).Error
	return err
}

// DeletePayOrderByIds 批量删除支付订单记录
// Author [yourname](https://github.com/yourname)
func (payOrderService *PayOrderService) DeletePayOrderByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]upay.PayOrder{}, "id in ?", IDs).Error
	return err
}

// UpdatePayOrder 更新支付订单记录
// Author [yourname](https://github.com/yourname)
func (payOrderService *PayOrderService) UpdatePayOrder(payOrder upay.PayOrder) (err error) {
	err = global.GVA_DB.Model(&upay.PayOrder{}).Where("id = ?", payOrder.ID).Updates(&payOrder).Error
	return err
}

// GetPayOrder 根据ID获取支付订单记录
// Author [yourname](https://github.com/yourname)
func (payOrderService *PayOrderService) GetPayOrder(ID string) (payOrder upay.PayOrder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&payOrder).Error
	return
}

// GetPayOrderInfoList 分页获取支付订单记录
// Author [yourname](https://github.com/yourname)
func (payOrderService *PayOrderService) GetPayOrderInfoList(info upayReq.PayOrderSearch) (list []upay.PayOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&upay.PayOrder{})
	var payOrders []upay.PayOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.Amount != nil {
		db = db.Where("amount <> ?", info.Amount)
	}
	if info.ActualAmount != nil {
		db = db.Where("actual_amount <> ?", info.ActualAmount)
	}
	if info.ChainType != "" {
		db = db.Where("chain_type = ?", info.ChainType)
	}
	if info.AppID != "" {
		db = db.Where("app_id = ?", info.AppID)
	}
	if info.Receiver != "" {
		db = db.Where("receiver = ?", info.Receiver)
	}
	if info.Hash != "" {
		db = db.Where("hash = ?", info.Hash)
	}
	if info.OrderNo != "" {
		db = db.Where("order_no = ?", info.OrderNo)
	}
	if info.MerchantOrderNo != "" {
		db = db.Where("merchant_order_no = ?", info.MerchantOrderNo)
	}
	if info.StartCompletedAt != nil && info.EndCompletedAt != nil {
		db = db.Where("completed_at BETWEEN ? AND ? ", info.StartCompletedAt, info.EndCompletedAt)
	}
	if info.RiskLevel != "" {
		db = db.Where("risk_level = ?", info.RiskLevel)
	}
	if info.UserID != 0 {
		db = db.Where("user_id = ?", info.UserID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["amount"] = true
	orderMap["actual_amount"] = true
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

	err = db.Find(&payOrders).Error
	return payOrders, total, err
}
func (payOrderService *PayOrderService) GetPayOrderDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	userID := make([]map[string]any, 0)

	global.GVA_DB.Table("sys_users").Select("username as label,id as value").Scan(&userID)
	res["userID"] = userID
	return
}
func (payOrderService *PayOrderService) GetPayOrderPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
