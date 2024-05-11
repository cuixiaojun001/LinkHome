package model

import (
	"encoding/json"
	"time"
)

const (
	// 订单状态
	NoPay    string = "no_pay"   // 未支付
	Payed    string = "payed"    // 已支付
	Ordered  string = "ordered"  // 已支付定金、已预订
	Canceled string = "canceled" // 已取消
	Finished string = "finished" // 订单已结束（合同结束）
	Deleted  string = "deleted"  // 已删除

	TenantID string = "tenant_id"
	HouseID  string = "house_id"
	State    string = "state"
)

type OrderModel struct {
	ID              int             `gorm:"column:id"`
	TradeNo         string          `gorm:"column:trade_no"`         // 最新的交易流水id
	TenantID        int             `gorm:"column:tenant_id"`        // 租客id
	LandlordID      int             `gorm:"column:landlord_id"`      // 房东id
	HouseID         int             `gorm:"column:house_id"`         // 房源id
	ContractContent string          `gorm:"column:contract_content"` // 合同内容
	State           string          `gorm:"column:state"`            // 订单状态
	PayMoney        int             `gorm:"column:pay_money;"`       // 支付总金额
	DepositFee      int             `gorm:"column:deposit_fee;"`     // 押金
	BargainMoney    float64         `gorm:"column:bargain_money;"`   // 房屋定金
	RentalDays      int             `gorm:"column:rental_days;"`     // 租赁天数
	StartDate       time.Time       `gorm:"column:start_date;"`      // 开始日期
	EndDate         time.Time       `gorm:"column:end_date;"`        // 结束日期
	JsonExtend      json.RawMessage `gorm:"column:json_extend;"`     // 扩展字段
	CreatedAt       time.Time       `gorm:"column:created_at"`
	UpdatedAt       time.Time       `gorm:"column:updated_at"`
}

func (o *OrderModel) TableName() string {
	return "user_order"
}
