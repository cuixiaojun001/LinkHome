package order

const (
	NoPay    string = "no_pay"   // 未支付
	Payed    string = "payed"    // 已支付
	Ordered  string = "ordered"  // 已支付定金、已预订
	Canceled string = "canceled" // 已取消
	Finished string = "finished" // 订单已结束（合同结束）
	Deleted  string = "deleted"  // 已删除
)
