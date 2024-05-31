package model

import (
	"encoding/json"
	"time"
)

const (
	Admin    string = "admin"
	Tenant   string = "tenant"
	Landlord string = "landlord"
	Steward  string = "steward"
)

type UserBasicInfo struct {
	ID         int             `gorm:"column:id"`
	Username   string          `gorm:"column:username"`
	Password   string          `gorm:"column:password"`
	Role       string          `gorm:"column:role"`
	Mobile     string          `gorm:"column:mobile"`
	State      string          `gorm:"column:state"`
	JsonExtend json.RawMessage `gorm:"column:json_extend"`
	CreatedAt  time.Time       `gorm:"column:created_at"`
	UpdatedAt  time.Time       `gorm:"column:updated_at"`
}

type UserProfileInfo struct {
	Id          int       `gorm:"column:id"`            // 主键id
	RealName    string    `gorm:"column:real_name"`     // 用户真实姓名
	Mobile      string    `gorm:"column:mobile"`        // 手机号
	Email       string    `gorm:"column:email"`         // 用户邮箱
	Avatar      string    `gorm:"column:avatar"`        // 用户头像
	IdCard      string    `gorm:"column:id_card"`       // 身份证号
	UserDesc    string    `gorm:"column:user_desc"`     // 用户简介
	Gender      string    `gorm:"column:gender"`        // 用户性别
	Hobby       string    `gorm:"column:hobby"`         // 用户爱好
	Career      string    `gorm:"column:career"`        // 职业
	IdCardFront string    `gorm:"column:id_card_front"` // 身份证正面
	IdCardBack  string    `gorm:"column:id_card_back"`  // 身份证反面
	AuthStatus  string    `gorm:"column:auth_status"`   // 实名认证状态
	State       string    `gorm:"column:state"`         // 用户状态
	JsonExtend  string    `gorm:"column:json_extend"`   // 扩展字段
	CreatedAt   time.Time `gorm:"column:created_at"`    // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at"`    // 更新时间
	AuthApplyAt time.Time `gorm:"column:auth_apply_at"` // 实名认证申请时间
}

type UserRentalDemandInfo struct {
	Id                   int             `gorm:"column:id"`                     // 主键id（需求id）
	UserID               int             `gorm:"column:user_id"`                // 用户id
	DemandTitle          string          `gorm:"column:demand_title"`           // 租房需求标题
	ExtendContent        string          `gorm:"column:extend_content"`         // 需求扩展内容
	City                 string          `gorm:"column:city"`                   // 城市
	RentTypeList         string          `gorm:"column:rent_type_list"`         // 租赁类型
	HouseTypeList        string          `gorm:"column:house_type_list"`        // 房屋类型
	HouseFacilities      string          `gorm:"column:house_facilities"`       // 房源设施要求
	TrafficInfoJson      json.RawMessage `gorm:"column:traffic_info_json"`      // 交通要求
	MinMoneyBudget       float64         `gorm:"column:min_money_budget"`       // 最低金额预算
	MaxMoneyBudget       float64         `gorm:"column:max_money_budget"`       // 最高金额预算
	Lighting             int             `gorm:"column:lighting"`               // 采光
	Floors               string          `gorm:"column:floors"`                 // 楼层需求
	CommutingTime        int             `gorm:"column:commuting_time"`         // 通勤时间（分钟）
	CompanyAddress       string          `gorm:"column:company_address"`        // 公司地址
	DesiredResidenceArea string          `gorm:"column:desired_residence_area"` // 期望居住地区
	Elevator             int             `gorm:"column:elevator"`               // 电梯要求
	State                string          `gorm:"column:state"`                  // 租房需求状态
	JsonExtend           json.RawMessage `gorm:"column:json_extend"`            // 扩展字段
	CreatedAt            time.Time       `gorm:"column:created_at"`             // 创建时间
	UpdatedAt            time.Time       `gorm:"column:updated_at"`             // 更新时间
}

func (u *UserBasicInfo) TableName() string {
	return "user_basic"
}

func (u *UserProfileInfo) TableName() string {
	return "user_profile"
}

func (u *UserRentalDemandInfo) TableName() string {
	return "user_rental_demand"
}
