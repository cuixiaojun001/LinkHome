package user

type TokenItem struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// LoginRequest 登录请求结构体
type LoginRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest 注册请求结构体
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	SmsCode  string `json:"sms_code" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

// PwdChangeRequest 修改密码请求结构体
type PwdChangeRequest struct {
	SrcPassword     string `json:"src_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type ProfileUpdateRequest struct {
	UserName   string `json:"username"`
	Mobile     string `json:"mobile"`
	RealName   string `json:"real_name"`
	Avatar     string `json:"avatar"`
	Mail       string `json:"mail"`
	IDCard     string `json:"id_card"`
	Gender     string `json:"gender"`
	Hobby      string `json:"hobby"`
	Career     string `json:"career"`
	AuthStatus string `json:"auth_status"`
}

// UserProfileItem 用户详细信息项
type UserProfileItem struct {
	UserID     string `json:"user_id"`     // 用户id
	Username   string `json:"username"`    // 用户名
	Mobile     string `json:"mobile"`      // 手机号
	Role       string `json:"role"`        // 用户角色
	State      string `json:"state"`       // 用户状态
	AuthStatus string `json:"auth_status"` // 实名认证状态

	AuthApplyTs int64  `json:"auth_apply_ts"` // 实名认证申请时间（时间戳）
	RealName    string `json:"real_name"`     // 用户真姓名
	Avatar      string `json:"avatar"`        // 用户头像
	Mail        string `json:"mail"`          // 电子邮件
	IDCard      string `json:"id_card"`       // 身份证号
	Gender      string `json:"gender"`        // 性别
	Hobby       string `json:"hobby"`         // 用户爱好
	Career      string `json:"career"`        // 用户职业
	IDCardFront string `json:"id_card_front"` // 身份证正面
	IDCardBack  string `json:"id_card_back"`  // 身份证反面
	CreateTs    int64  `json:"create_ts"`     // 用户创建时间（时间戳）
}

// RentalDemandRequest 租房需求发布入参
type RentalDemandRequest struct {
	ID                   int                    `json:"id"`                                  // 主键ID
	DemandTitle          string                 `json:"demand_title" binding:"required"`     // 租房需求标题
	City                 string                 `json:"city" binding:"required"`             // 期望城市
	MinMoneyBudget       float64                `json:"min_money_budget" binding:"required"` // 最低金额预算
	MaxMoneyBudget       float64                `json:"max_money_budget" binding:"required"` // 最高金额预算
	DesiredResidenceArea string                 `json:"desired_residence_area"`              // 期望居住地区
	TrafficInfoJson      map[string]interface{} `json:"traffic_info_json"`                   // 交通要求
	HouseFacilities      []int                  `json:"house_facilities"`                    // 房源设施要求
	Floors               []int                  `json:"floors"`                              // 房屋楼层要求

	RentTypeList  []string `json:"rent_type_list"`  // 租赁类型
	HouseTypeList []string `json:"house_type_list"` // 房源类型
	Lighting      int      `json:"lighting"`        // 采光要求 0:差 1:一般 2:正常 3:良好 4:极好
	Elevator      int      `json:"elevator"`        // 电梯要求 0:不需要 1:需要 2:无要求

	CommutingTime  int    `json:"commuting_time"`  // 通勤时间
	CompanyAddress string `json:"company_address"` // 公司地址
	ExtendContent  string `json:"extend_content"`  // 租房需求扩展内容
}

// UserRealNameAuthRequest 用户实名认证请求结构体
type UserRealNameAuthRequest struct {
	UserID      int    `json:"user_id"`       // 用户id
	RealName    string `json:"real_name"`     // 真实姓名
	IDCard      string `json:"id_card"`       // 身份证号
	IDCardFront string `json:"id_card_front"` // 身份证正面
	IDCardBack  string `json:"id_card_back"`  // 身份证反面
}

// UserRealNameAuthResponse 用户实名认证响应结构体
type UserRealNameAuthResponse struct {
	UserID      int    `json:"user_id"`       // 用户id
	State       string `json:"state"`         // 认证状态
	AuthStatus  string `json:"auth_status"`   // 实名认证状态
	RealName    string `json:"real_name"`     // 真实姓名
	IDCard      string `json:"id_card"`       // 身份证号
	IDCardFront string `json:"id_card_front"` // 身份证正面
	IDCardBack  string `json:"id_card_back"`  // 身份证反面
}

type RentalDemandListRequest struct {
	QueryParams struct {
		UserID int `json:"user_id"` // 用户id
	} `json:"query_params"` // QueryParams 房源列表查询参数
	Offset    int      `json:"offset"`    // Offset 分页偏移量
	Limit     int      `json:"limit"`     // Limit 每页显示数量
	Orderings []string `json:"orderings"` // Orderings 排序字段
}

type RentalDemandListResponse struct {
	Total      int                    `json:"total"`       // Total 总数
	HasMore    bool                   `json:"has_more"`    // HasMore 是否有下一页
	NextOffset int                    `json:"next_offset"` // NextOffset offset下次起步
	DataList   []RentalDemandListItem `json:"data_list"`   // DataList 用户列表
}

type RentalDemandListItem struct {
	ID             int     `json:"id"`                                  // 主键ID
	UserID         int     `json:"user_id"`                             // 用户ID
	DemandTitle    string  `json:"demand_title" binding:"required"`     // 租房需求标题
	City           string  `json:"city" binding:"required"`             // 期望城市
	MinMoneyBudget float64 `json:"min_money_budget" binding:"required"` // 最低金额预算
	MaxMoneyBudget float64 `json:"max_money_budget" binding:"required"` // 最高金额预算

	RentTypeList    []string `json:"rent_type_list"`   // 租赁类型
	HouseTypeList   []string `json:"house_type_list"`  // 房源类型
	HouseFacilities []int    `json:"house_facilities"` // 房源设施要求
	Floors          []int    `json:"floors"`           // 房屋楼层要求
	CommutingTime   int      `json:"commuting_time"`   // 通勤时间
	CompanyAddress  string   `json:"company_address"`  // 公司地址

	Lighting int    `json:"lighting"` // 采光要求 0:差 1:一般 2:正常 3:良好 4:极好
	Elevator int    `json:"elevator"` // 电梯要求 0:不需要 1:需要 2:无要求
	State    string `json:"state"`    // 租房需求状态 0:未发布 1:已发布 2:已删除

	DesiredResidenceArea string `json:"desired_residence_area"` // 期望居住地区
	ExtendContent        string `json:"extend_content"`         // 租房需求扩展内容
	CreateTs             int64  `json:"create_ts"`              // 用户创建时间（时间戳）
}

type RentalDemandDetailResponse struct {
	RentalDemandListItem
	UserInfo UserItem `json:"user_info"` // 用户信息
}

type UserItem struct {
	ID       int    `json:"id"`       // 用户ID
	UserName string `json:"username"` // 用户名
	Mobile   string `json:"mobile"`   // 手机号
	Role     string `json:"role"`     // 用户角色
}
