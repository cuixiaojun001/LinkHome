package model

import "time"

type PaymentTrade struct {
	ID          int       `gorm:"id"`           // 主键id
	OrderId     int       `gorm:"order_id"`     // 订单id
	UserId      int       `gorm:"user_id"`      // 用户id（记录谁付的钱）
	TradeNo     string    `gorm:"trade_no"`     // 订单交易流水号
	Scene       string    `gorm:"scene"`        // 交易场景
	TransAmount int       `gorm:"trans_amount"` // 交易金额
	CreatedAt   time.Time `gorm:"created_at"`   // 创建时间
	UpdatedAt   time.Time `gorm:"updated_at"`   // 更新时间
}

// TableName 表名称
func (*PaymentTrade) TableName() string {
	return "payment_trade"
}
