package admin

import (
	"github.com/cuixiaojun001/LinkHome/library/orm"
	"github.com/cuixiaojun001/LinkHome/services/house"
	"github.com/cuixiaojun001/LinkHome/services/order"
	"github.com/cuixiaojun001/LinkHome/services/user"
)

type UserListRequest struct {
	QueryParams struct {
		UserID     int    `json:"user_id"`     // 用户id
		Mobile     string `json:"mobile"`      // 手机号
		RealName   string `json:"real_name"`   // 用户真实姓名
		Gender     string `json:"gender"`      // 性别
		Career     string `json:"career"`      // 用户职业
		State      string `json:"state"`       // 租赁类型
		AuthStatus string `json:"auth_status"` // 用户实名认证状态
	} `json:"query_params"` // QueryParams 房源列表查询参数
	Offset    int      `json:"offset"`    // Offset 分页偏移量
	Limit     int      `json:"limit"`     // Limit 每页显示数量
	Orderings []string `json:"orderings"` // Orderings 排序字段
}

func (r *UserListRequest) GenQuery() orm.IQuery {
	query := orm.NewQuery()
	if r.QueryParams.UserID > 0 {
		query.ExactMatch("id", r.QueryParams.UserID)
	}
	if r.QueryParams.Mobile != "" {
		query.ExactMatch("mobile", r.QueryParams.Mobile)
	}
	if r.QueryParams.RealName != "" {
		query.ExactMatch("real_name", r.QueryParams.RealName)
	}
	if r.QueryParams.Gender != "" {
		query.ExactMatch("gender", r.QueryParams.Gender)
	}
	if r.QueryParams.Career != "" {
		query.ExactMatch("career", r.QueryParams.Career)
	}
	if r.QueryParams.State != "" {
		query.ExactMatch("state", r.QueryParams.State)
	}
	if r.QueryParams.AuthStatus != "" {
		query.ExactMatch("auth_status", r.QueryParams.AuthStatus)
	}
	if r.Limit > 0 && r.Offset >= 0 {
		query.SetPagination(r.Offset, r.Limit)
	}
	return query
}

type UserListResponse struct {
	Total      int                    `json:"total"`       // Total 总数
	HasMore    bool                   `json:"has_more"`    // HasMore 是否有下一页
	NextOffset int                    `json:"next_offset"` // NextOffset offset下次起步
	DataList   []user.UserProfileItem `json:"data_list"`   // DataList 用户列表
}

type OrderListRequest struct {
	QueryParams struct {
		OrderID    int    `json:"order_id"`    // 订单id
		TenantID   int    `json:"tenant_id"`   // 租客id
		LandlordID int    `json:"landlord_id"` // 房东id
		HouseID    int    `json:"house_id"`    // 房源id
		StartDate  int    `json:"start_date"`  // 开始日期
		EndData    int    `json:"end_data"`    // 结束日期
		PayMoney   string `json:"pay_money"`   // 付款金额
		DepositFee string `json:"deposit_fee"` // 押金金额
		RentalDays int    `json:"rental_days"` // 租赁天数
		State      string `json:"state"`       // 订单状态
	} `json:"query_params"` // QueryParams 房源列表查询参数
	Offset    int      `json:"offset"`    // Offset 分页偏移量
	Limit     int      `json:"limit"`     // Limit 每页显示数量
	Orderings []string `json:"orderings"` // Orderings 排序字段
}

func (r *OrderListRequest) GenQuery() orm.IQuery {
	query := orm.NewQuery()
	if r.QueryParams.OrderID > 0 {
		query.ExactMatch("id", r.QueryParams.OrderID)
	}
	if r.QueryParams.TenantID > 0 {
		query.ExactMatch("tenant_id", r.QueryParams.TenantID)
	}
	if r.QueryParams.LandlordID > 0 {
		query.ExactMatch("landlord_id", r.QueryParams.LandlordID)
	}
	if r.QueryParams.HouseID > 0 {
		query.ExactMatch("house_id", r.QueryParams.HouseID)
	}
	if r.QueryParams.StartDate > 0 {
		query.ExactMatch("start_date", r.QueryParams.StartDate)
	}
	if r.QueryParams.EndData > 0 {
		query.ExactMatch("end_data", r.QueryParams.EndData)
	}
	if r.QueryParams.PayMoney != "" {
		query.ExactMatch("pay_money", r.QueryParams.PayMoney)
	}
	if r.QueryParams.DepositFee != "" {
		query.ExactMatch("deposit_fee", r.QueryParams.DepositFee)
	}
	if r.QueryParams.RentalDays > 0 {
		query.ExactMatch("rental_days", r.QueryParams.RentalDays)
	}
	if r.QueryParams.State != "" {
		query.ExactMatch("state", r.QueryParams.State)
	}
	if r.Limit > 0 && r.Offset >= 0 {
		query.SetPagination(r.Offset, r.Limit)
	}
	return query

}

type OrderListResponse struct {
	Total      int                 `json:"total"`       // Total 总数
	HasMore    bool                `json:"has_more"`    // HasMore 是否有下一页
	NextOffset int                 `json:"next_offset"` // NextOffset offset下次起步
	DataList   []UserOrderListItem `json:"data_list"`   // DataList 用户列表
}

type UserOrderListItem struct {
	OrderID         int                `json:"order_id"`         // 订单id
	TenantID        int                `json:"tenant_id"`        // 租客id
	LandlordID      int                `json:"landlord_id"`      // 房东id
	HouseID         int                `json:"house_id"`         // 房源id
	StartDate       string             `json:"start_date"`       // 开始日期
	EndDate         string             `json:"end_date"`         // 结束日期
	State           string             `json:"state"`            // 订单状态
	ContractContent string             `json:"contract_content"` // 合同内容
	PayMoney        string             `json:"pay_money"`        // 支付金额
	BargainMoney    string             `json:"bargain_money"`    // 房屋预定金
	DepositFee      string             `json:"deposit_fee"`      // 押金
	RentalDays      int                `json:"rental_days"`      // 租赁天数
	UserInfo        order.UserInfoItem `json:"user_info"`        // 租客信息
	LandlordInfo    order.UserInfoItem `json:"landlord_info"`    // 房东信息
	HouseInfo       house.HouseDetail  `json:"house_info"`       // 房源信息
	CreateTs        int64              `json:"create_ts"`        // 创建时间
	UpdateTs        int64              `json:"update_ts"`        // 更新时间
}

type UserInfoItem struct {
	UserID   int    `json:"user_id"`   // 用户id
	RearName string `json:"real_name"` // 真实姓名
	Mobile   string `json:"mobile"`    // 手机号
}
