// 自动生成模板PayOrder
package upay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// 支付订单 结构体  PayOrder
type PayOrder struct {
	global.GVA_MODEL
	AppID           string          `json:"appId" form:"appId" gorm:"column:app_id;comment:应用标识;" binding:"required"`                                             //应用标识
	MerchantOrderNo string          `json:"merchantOrderNo" form:"merchantOrderNo" gorm:"index;column:merchant_order_no;comment:商户订单号;" binding:"required"`       //商户订单号
	ChainType       string          `json:"chainType" form:"chainType" gorm:"default:1;column:chain_type;comment:区块链类型;" binding:"required"`                      //区块链类型
	Amount          decimal.Decimal `json:"crypto" form:"crypto" gorm:"column:amount;comment:金额;size:DECIMAL(10, 4);" binding:"required"`                         //金额
	Attach          string          `json:"attach" form:"attach" gorm:"column:attach;comment:用户自定义数据;"`                                                           //用户自定义数据
	RedirectUrl     string          `json:"redirectUrl" form:"redirectUrl" gorm:"default:0;column:redirect_url;comment:支付成功后，前端重定向地址;"`                           //重定向地址
	NotifyUrl       string          `json:"notifyUrl" form:"notifyUrl" gorm:"default:0;column:notify_url;comment:接收异步通知的回调地址;"`                                   //回调地址
	OrderNo         string          `json:"orderNo" form:"orderNo" gorm:"uniqueIndex;column:order_no;comment:订单号;"`                                               //订单号
	Status          string          `json:"status" form:"status" gorm:"default:0;column:status;comment:支付状态;"`                                                    //支付状态
	ActualAmount    decimal.Decimal `json:"actualCrypto" form:"actualCrypto" gorm:"column:actual_amount;default:0;comment:实际金额;size:DECIMAL(10, 4);"`             //实际金额
	Fee             decimal.Decimal `json:"poundage" form:"poundage" gorm:"default:0;column:fee;default:0;comment:手续费;size:DECIMAL(10, 4);"`                      //手续费
	ActualFee       decimal.Decimal `json:"actualPoundage" form:"actualPoundage" gorm:"default:0;column:actual_fee;default:0;comment:实际手续费;size:DECIMAL(10, 4);"` //实际手续费
	Receiver        string          `json:"receiver" form:"receiver" gorm:"column:receiver;comment:收款地址;" binding:"required"`                                     //收款地址
	Hash            string          `json:"hash" form:"hash" gorm:"column:hash;comment:交易哈希;"`                                                                    //交易哈希
	ExpiredAt       time.Time       `json:"expiredAt" form:"expiredAt" gorm:"column:expired_at;comment:完成时间;"`                                                    //过期时间
	CompletedAt     *time.Time      `json:"completedAt" form:"completedAt" gorm:"column:completed_at;comment:完成时间;"`                                              //完成时间
	RiskLevel       string          `json:"riskLevel" form:"riskLevel" gorm:"default:0;column:risk_level;comment:风控级别;"`                                          //风控级别
	UserID          uint            `json:"userID" form:"userID" gorm:"column:user_id;comment:用户;"`                                                               //用户
}

// TableName 支付订单 PayOrder自定义表名 upay_pay_order
func (PayOrder) TableName() string {
	return "upay_pay_order"
}

func (order *PayOrder) BeforeCreate(tx *gorm.DB) (err error) {
	increment := decimal.NewFromFloat(0.0001)
	upperLimit := order.Amount.Add(decimal.NewFromFloat(0.0010))
	amount, err := getNextAmount(tx, order.Amount, increment, upperLimit)
	if err != nil {
		return err
	}
	order.Amount = amount

	if len(order.OrderNo) == 0 {
		order.OrderNo = uuid.New().String()
	}

	wallet, err := getAddressByAppID(tx, order.UserID, order.ChainType)
	if err != nil {
		return err
	}
	order.Receiver = wallet.Address

	order.ExpiredAt = time.Now().Add(time.Minute * 30)
	return nil
}

func getAddressByAppID(db *gorm.DB, userId uint, chainType string) (*WalletAddress, error) {
	var wallet WalletAddress

	if err := db.First(&wallet, "user_id = ? AND chain_type = ? AND status = ?", userId, chainType, 0).Error; err != nil {
		return nil, err
	}

	if err := db.Where("user_id = ? AND chain_type = ? AND status = ?", userId, chainType, 0).
		Order("RAND()").
		Limit(1).
		Find(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}
