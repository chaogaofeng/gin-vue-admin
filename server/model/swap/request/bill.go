package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BillSearch struct {
	StartCreatedAt  *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt    *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	No              string     `json:"no" form:"no" `
	OrderNo         string     `json:"orderNo" form:"orderNo" `
	UserID          *int       `json:"user_id" form:"user_id" `
	PlanID          *int       `json:"plan_id" form:"plan_id" `
	Hash            string     `json:"hash" form:"hash" `
	Status          string     `json:"status" form:"status" `
	StartCompleteAt *time.Time `json:"startCompleteAt" form:"startCompleteAt"`
	EndCompleteAt   *time.Time `json:"endCompleteAt" form:"endCompleteAt"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
