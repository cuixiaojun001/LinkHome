package order

type CreateOrderRequest struct {
	HouseID   int    `json:"house_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type UserOrderListResponse struct {
	UserOrders []UserOrderListItem `json:"user_orders"` // 用户订单列表
}

type UserOrderListItem struct {
	OrderID         int           `json:"order_id"`         // 订单id
	TenantID        int           `json:"tenant_id"`        // 租客id
	LandlordID      int           `json:"landlord_id"`      // 房东id
	HouseID         int           `json:"house_id"`         // 房源id
	StartDate       string        `json:"start_date"`       // 开始日期
	EndDate         string        `json:"end_date"`         // 结束日期
	State           string        `json:"state"`            // 订单状态
	ContractContent string        `json:"contract_content"` // 合同内容
	PayMoney        string        `json:"pay_money"`        // 支付金额
	BargainMoney    string        `json:"bargain_money"`    // 房屋预定金
	DepositFee      string        `json:"deposit_fee"`      // 押金
	RentalDays      int           `json:"rental_days"`      // 租赁天数
	UserInfo        UserInfoItem  `json:"user_info"`        // 租客信息
	LandlordInfo    UserInfoItem  `json:"landlord_info"`    // 房东信息
	HouseInfo       HouseInfoItem `json:"house_info"`       // 房源信息
	CreateTs        int64         `json:"create_ts"`        // 创建时间
	UpdateTs        int64         `json:"update_ts"`        // 更新时间
}

type UserInfoItem struct {
	UserID   int    `json:"user_id"`   // 用户id
	RearName string `json:"real_name"` // 真实姓名
	Mobile   string `json:"mobile"`    // 手机号
}

type HouseInfoItem struct {
	HouseID      int    `json:"house_id"`      // 房源id
	Title        string `json:"title"`         // 房源标题
	Address      string `json:"address"`       // 房源地址
	IndexImg     string `json:"index_img"`     // 房源图片
	RentType     string `json:"rent_type"`     // 房源类型
	RentMoney    string `json:"rent_money"`    // 房源租金
	StrataFee    string `json:"strata_fee"`    // 房源管理费
	DepositRatio string `json:"deposit_ratio"` // 房源租金扣押比率
	PayRatio     string `json:"pay_ratio"`     // 房源租金支付比率
}
