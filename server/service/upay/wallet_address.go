package upay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/upay"
	upayReq "github.com/flipped-aurora/gin-vue-admin/server/model/upay/request"
)

type WalletAddressService struct{}

// CreateWalletAddress 创建收款钱包记录
// Author [yourname](https://github.com/yourname)
func (walletAddressService *WalletAddressService) CreateWalletAddress(walletAddress *upay.WalletAddress) (err error) {
	err = global.GVA_DB.Create(walletAddress).Error
	return err
}

// DeleteWalletAddress 删除收款钱包记录
// Author [yourname](https://github.com/yourname)
func (walletAddressService *WalletAddressService) DeleteWalletAddress(ID string) (err error) {
	err = global.GVA_DB.Delete(&upay.WalletAddress{}, "id = ?", ID).Error
	return err
}

// DeleteWalletAddressByIds 批量删除收款钱包记录
// Author [yourname](https://github.com/yourname)
func (walletAddressService *WalletAddressService) DeleteWalletAddressByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]upay.WalletAddress{}, "id in ?", IDs).Error
	return err
}

// UpdateWalletAddress 更新收款钱包记录
// Author [yourname](https://github.com/yourname)
func (walletAddressService *WalletAddressService) UpdateWalletAddress(walletAddress upay.WalletAddress) (err error) {
	err = global.GVA_DB.Model(&upay.WalletAddress{}).Where("id = ?", walletAddress.ID).Updates(&walletAddress).Error
	return err
}

// GetWalletAddress 根据ID获取收款钱包记录
// Author [yourname](https://github.com/yourname)
func (walletAddressService *WalletAddressService) GetWalletAddress(ID string) (walletAddress upay.WalletAddress, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&walletAddress).Error
	return
}

// GetWalletAddressInfoList 分页获取收款钱包记录
// Author [yourname](https://github.com/yourname)
func (walletAddressService *WalletAddressService) GetWalletAddressInfoList(info upayReq.WalletAddressSearch) (list []upay.WalletAddress, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&upay.WalletAddress{})
	var walletAddresss []upay.WalletAddress
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Address != "" {
		db = db.Where("address = ?", info.Address)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.ChainType != "" {
		db = db.Where("chain_type = ?", info.ChainType)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.UserID != 0 {
		db = db.Where("user_id = ?", info.UserID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&walletAddresss).Error
	return walletAddresss, total, err
}
func (walletAddressService *WalletAddressService) GetWalletAddressDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	userID := make([]map[string]any, 0)

	global.GVA_DB.Table("sys_users").Select("username as label,id as value").Scan(&userID)
	res["userID"] = userID
	return
}
func (walletAddressService *WalletAddressService) GetWalletAddressPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
