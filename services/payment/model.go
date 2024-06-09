package payment

type OrderPaymentRequest struct {
	StartDate string `json:"start_date"` // 开始入住日期
	EndDate   string `json:"end_date"`   // 结束入住日期
	PayScene  string `json:"pay_scene"`  // 支付场景 full_payment:全额支付, deposit_payment:定金支付 balance_payment:尾款支付
}

type OrderPaymentResponse struct {
	OrderID   int    `json:"order_id"`   // 订单ID
	AliPayURL string `json:"alipay_url"` // 阿里支付url
}
