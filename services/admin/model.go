package admin

import (
	"github.com/cuixiaojun001/LinkHome/library/orm"
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
