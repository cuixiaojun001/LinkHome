package user

import "github.com/cuixiaojun001/linkhome/modules/house/model"

type TokenItem struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Role string

const (
	Admin    Role = "admin"
	Tenant   Role = "tenant"
	Landlord Role = "landlord"
	Steward  Role = "steward"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	SmsCode  string `json:"sms_code" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     Role   `json:"role" binding:"required"`
}

type PwdChangeRequest struct {
	SrcPassword     string `json:"src_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type UserAuthStatus string

const (
	unauthorized UserAuthStatus = "unauthorized" //  未实名认证
	authorized   UserAuthStatus = "authorized"   //  已实名认证
	auditing     UserAuthStatus = "auditing"     //  审核中
	unapprove    UserAuthStatus = "unapprove"    //  审核未通过
)

type ProfileUpdateRequest struct {
	UserName   string         `json:"username"`
	Mobile     string         `json:"mobile"`
	RealName   string         `json:"real_name"`
	Avatar     string         `json:"avatar"`
	Mail       string         `json:"mail"`
	IDCard     string         `json:"id_card"`
	Gender     string         `json:"gender"`
	Hobby      string         `json:"hobby"`
	Career     string         `json:"career"`
	AuthStatus UserAuthStatus `json:"auth_status"`
}

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

	RentTypeList  []model.RentType        `json:"rent_type_list"`  // 租赁类型
	HouseTypeList []model.HouseType       `json:"house_type_list"` // 房源类型
	Lighting      HouseLightingEnum       `json:"lighting"`        // 采光要求
	Elevator      HouseElevatorDemandEnum `json:"elevator"`        // 电梯要求

	CommutingTime  int    `json:"commuting_time"`  // 通勤时间
	CompanyAddress string `json:"company_address"` // 公司地址
	ExtendContent  string `json:"extend_content"`  // 租房需求扩展内容
}

type HouseLightingEnum int

const (
	bad       HouseLightingEnum = iota + 0 // 差
	general                                // 一般
	normal                                 // 正常
	good                                   // 良好
	excellent                              // 极好

)

type HouseElevatorDemandEnum int

const (
	not_Required   HouseElevatorDemandEnum = iota + 0 // 不需要
	required                                          // 需要
	no_requirement                                    // 无要求
)
