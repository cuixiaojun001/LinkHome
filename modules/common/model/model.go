package model

import (
	"encoding/json"
	"time"
)

type District struct {
	ID       int    `gorm:"id" json:"id"`
	Name     string `gorm:"name" json:"name"`
	ParentID int    `gorm:"parent_id" json:"parent_id"`
}

func (c *District) TableName() string {
	return "areas"
}

type City struct {
	ID           int        `gorm:"id" json:"id"`
	Name         string     `gorm:"name" json:"name"`
	ParentID     int        `gorm:"parent_id" json:"parent_id"`
	DistrictList []District `gorm:"-" json:"district_list"`
}

func (c *City) TableName() string {
	return "areas"
}

type Province struct {
	ID       int    `gorm:"id" json:"id"`
	Name     string `gorm:"name" json:"name"`
	ParentId int    `gorm:"parent_id" json:"-"`
	CityList []City `gorm:"-" json:"city_list"`
}

func (c *Province) TableName() string {
	return "areas"
}

type ContractTemplate struct {
	ID              int             `gorm:"id"`               // 主键id（合同模板id）
	TemplateContent string          `gorm:"template_content"` // 合同模板内容
	RenderParams    json.RawMessage `gorm:"render_params"`    // 渲染参数
	ApiParams       json.RawMessage `gorm:"api_params"`       // 需要的接口参数
	State           string          `gorm:"state"`            // 状态
	JsonExtend      json.RawMessage `gorm:"json_extend"`      // 扩展字段
	CreatedAt       time.Time       `gorm:"created_at"`
	UpdatedAt       time.Time       `gorm:"updated_at"`
}

func (c *ContractTemplate) TableName() string {
	return "template"
}

type FieldInfo struct {
	FieldName      string `json:"field_name"`
	DBModelManager string `json:"db_model_manager"`
	Role           string `json:"role,omitempty"`
}

type RenderParams map[string]FieldInfo
