package model

import (
	"encoding/json"
	"time"
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
	UpdatedAt  time.Time       `gorm:"column:update_at"`
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
	UpdatedAt   time.Time `gorm:"column:update_at"`     // 更新时间
	AuthApplyAt time.Time `gorm:"column:auth_apply_at"` // 实名认证申请时间
}

func (u *UserBasicInfo) TableName() string {
	return "user_basic"
}

func (u *UserProfileInfo) TableName() string {
	return "user_profile"
}
