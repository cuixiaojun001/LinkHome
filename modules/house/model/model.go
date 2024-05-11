package model

import (
	"encoding/json"
	"time"
)

const (
	Rent    string = "rent"     // rent 已出租
	NotRent string = "not_rent" // not_rent 未出租
	Ordered string = "ordered"  // ordered 已预定

	Up       string = "up"       // up 房源状态: 已上架
	Down     string = "down"     // down 房源状态: 已下架
	Auditing string = "auditing" // auditing 房源状态: 审核中
	Deleted  string = "deleted"  // deleted 房源状态: 已删除

	Whole string = "whole" // whole 整租
	Share string = "share" // share 合租
)

type HouseInfo struct {
	ID              int       `gorm:"column:id" json:"house_id"`                     // 主键ID
	RentType        string    `gorm:"column:rent_type" json:"rent_type"`             // 出租类型, 整租 合租 日租
	HouseType       string    `gorm:"column:house_type" json:"house_type"`           // 房屋类型，小区房、公寓、自家房
	Title           string    `gorm:"column:title" json:"title"`                     // 房屋标题
	IndexImg        string    `gorm:"column:index_img" json:"index_img"`             // 房屋首页图片
	HouseDesc       string    `gorm:"column:house_desc" json:"house_desc"`           // 房屋描述
	City            string    `gorm:"column:city" json:"city"`                       // 房屋所在城市
	District        string    `gorm:"column:district" json:"district"`               // 区县
	Address         string    `gorm:"column:address" json:"address"`                 // 房屋地址
	RentState       string    `gorm:"column:rent_state" json:"rent_state"`           // 出租状态
	RentMoney       int       `gorm:"column:rent_money" json:"rent_money"`           // 租赁金额（单位/分、元/100）
	BargainMoney    float64   `gorm:"column:bargain_money"`                          // 定金
	RentTimeUnit    string    `gorm:"column:rent_time_unit"`                         // 租赁时间单位，默认month（月结）
	WaterRent       float64   `gorm:"column:water_rent"`                             // 水费（单位/分，元/100）
	ElectricityRent float64   `gorm:"column:electricity_rent"`                       // 电费（单位/分，元/100）
	StrataFee       int       `gorm:"column:strata_fee"`                             // 管理费（单位/分，元/100）
	DepositRatio    int       `gorm:"column:deposit_ratio"`                          // 租赁费用的押金倍数（押几付几）
	PayRatio        int       `gorm:"column:pay_ratio"`                              // 租赁费用的支付倍数（押几付几）
	BedroomNum      int       `gorm:"column:bedroom_num" json:"bedroom_num"`         // 卧室数量
	LivingRoomNum   int       `gorm:"column:living_room_num" json:"living_room_num"` // 客厅数量
	KitchenNum      int       `gorm:"column:kitchen_num" json:"kitchen_num"`         // 厨房数量
	ToiletNum       int       `gorm:"column:toilet_num" json:"toilet_num"`           // 卫生间数量
	Area            int       `gorm:"column:area" json:"area"`                       // 房屋总体面积
	PublishedAt     time.Time `gorm:"column:published_at" json:"-"`                  // 发布时间
	State           string    `gorm:"column:state" json:"state"`                     // 房屋状态
	JsonExtend      *string   `gorm:"column:json_extend"`                            // 扩展字段
	CreatedAt       time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"-"`
	HouseOwner      int       `gorm:"column:house_owner"` // 房屋拥有者
}

func (t *HouseInfo) TableName() string {
	return "house_info"
}

type HouseDetailInfo struct {
	Id              int             `gorm:"column:id"`                                         // 主键id
	HouseID         int             `gorm:"column:house_id"`                                   // 房屋id
	HouseOwner      int             `gorm:"column:house_owner" json:"house_owner"`             // 房源拥有者
	ContactId       int             `gorm:"column:contact_id" json:"contact_id"`               // 房源联系人id
	Address         string          `gorm:"column:address" json:"address"`                     // 房屋详细地址
	RoomNum         int             `gorm:"column:room_num" json:"room_num"`                   // 房间号
	DisplayContent  *string         `gorm:"column:display_content" json:"display_content"`     // 房屋展示内容json
	Floor           int             `gorm:"column:floor" json:"floor"`                         // 房屋所在楼层
	MaxFloor        int             `gorm:"column:max_floor" json:"max_floor"`                 // 房屋最大楼层
	HasElevator     int8            `gorm:"column:has_elevator" json:"has_elevator"`           // 是否有电梯（0没有、1有)
	BuildYear       string          `gorm:"column:build_year" json:"build_year"`               // 建成年份
	Direction       string          `gorm:"column:direction" json:"direction"`                 // 房屋朝向
	Lighting        int8            `gorm:"column:lighting" json:"lighting"`                   // 房源采光情况
	NearTrafficJson string          `gorm:"column:near_traffic_json" json:"near_traffic_json"` // 附近交通信息
	CertificateNo   string          `gorm:"column:certificate_no" json:"certificate_no"`       // 房产证号
	LocationInfo    json.RawMessage `gorm:"column:location_info" json:"location_info"`         // 房源地理位置信息
	JsonExtend      *string         `gorm:"column:json_extend"`                                // 扩展字段
	CreatedAt       time.Time       `gorm:"column:created_at"`                                 // 创建时间
	UpdatedAt       time.Time       `gorm:"column:updated_at"`                                 // 更新时间
}

func (t *HouseDetailInfo) TableName() string {
	return "house_detail"
}

type Facility struct {
	ID        int       `gorm:"column:id" json:"facility_id"`
	Name      string    `gorm:"column:name" json:"name"`
	Icon      string    `gorm:"column:icon" json:"icon"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

func (f *Facility) TableName() string {
	return "house_facility"
}

type HouseFacilityMapping struct {
	ID         int       `gorm:"column:id"`
	HouseID    int       `gorm:"column:house_id"`
	FacilityID int       `gorm:"column:facility_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (f *HouseFacilityMapping) TableName() string {
	return "house_facility_mapping"
}
