package swap

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/swap"
	swapReq "github.com/flipped-aurora/gin-vue-admin/server/model/swap/request"
	"gorm.io/gorm"
)

type BillService struct{}

// CreateBill 创建账单记录
// Author [yourname](https://github.com/yourname)
func (billService *BillService) CreateBill(bill *swap.Bill) (err error) {
	var item swap.Bill
	if !errors.Is(global.GVA_DB.Where("no = ?", bill.No).First(&item).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("订单号已存在")
	}

	err = global.GVA_DB.Create(bill).Error
	return err
}

// DeleteBill 删除账单记录
// Author [yourname](https://github.com/yourname)
func (billService *BillService) DeleteBill(ID string) (err error) {
	err = global.GVA_DB.Delete(&swap.Bill{}, "id = ?", ID).Error
	return err
}

// DeleteBillByIds 批量删除账单记录
// Author [yourname](https://github.com/yourname)
func (billService *BillService) DeleteBillByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]swap.Bill{}, "id in ?", IDs).Error
	return err
}

// UpdateBill 更新账单记录
// Author [yourname](https://github.com/yourname)
func (billService *BillService) UpdateBill(bill swap.Bill) (err error) {
	err = global.GVA_DB.Model(&swap.Bill{}).Where("id = ?", bill.ID).Updates(&bill).Error
	return err
}

// GetBill 根据ID获取账单记录
// Author [yourname](https://github.com/yourname)
func (billService *BillService) GetBill(ID string) (bill swap.Bill, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&bill).Error
	return
}
func (billService *BillService) GetBillByNo(no string) (bill swap.Bill, err error) {
	err = global.GVA_DB.Where("no = ?", no).First(&bill).Error
	return
}

// GetBillInfoList 分页获取账单记录
// Author [yourname](https://github.com/yourname)
func (billService *BillService) GetBillInfoList(info swapReq.BillSearch) (list []swap.Bill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&swap.Bill{})
	var bills []swap.Bill
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.No != "" {
		db = db.Where("no LIKE ?", "%"+info.No+"%")
	}
	if info.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+info.OrderNo+"%")
	}
	if info.PlanID != nil {
		db = db.Where("plan_id = ?", info.PlanID)
	}
	if info.UserID != nil {
		db = db.Where("user_id = ?", info.UserID)
	}
	if info.Hash != "" {
		db = db.Where("hash = ?", info.Hash)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.StartCompleteAt != nil && info.EndCompleteAt != nil {
		db = db.Where("complete_at BETWEEN ? AND ? ", info.StartCompleteAt, info.EndCompleteAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["plan_id"] = true
	orderMap["count"] = true
	orderMap["status"] = true
	orderMap["amount"] = true
	orderMap["complete_at"] = true
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

	err = db.Find(&bills).Error
	return bills, total, err
}
func (billService *BillService) GetBillDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	plan_id := make([]map[string]any, 0)

	global.GVA_DB.Table("billing_plan").Select("name as label,id as value").Scan(&plan_id)
	res["plan_id"] = plan_id
	return
}
func (billService *BillService) GetBillPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
