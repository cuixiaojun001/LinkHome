package model

import (
	"fmt"
	"github.com/cuixiaojun001/linkhome/common/logger"
	"strings"
	"time"
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
	BargainMoney    int       `gorm:"column:bargain_money"`                          // 定金
	RentTimeUnit    string    `gorm:"column:rent_time_unit"`                         // 租赁时间单位，默认month（月结）
	WaterRent       int       `gorm:"column:water_rent"`                             // 水费（单位/分，元/100）
	ElectricityRent int       `gorm:"column:electricity_rent"`                       // 电费（单位/分，元/100）
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
	Id              int       `gorm:"column:id"`                                         // 主键id
	HouseID         int       `gorm:"column:house_id"`                                   // 房屋id
	HouseOwner      int       `gorm:"column:house_owner" json:"house_owner"`             // 房源拥有者
	ContactId       int       `gorm:"column:contact_id" json:"contact_id"`               // 房源联系人id
	Address         string    `gorm:"column:address" json:"address"`                     // 房屋详细地址
	RoomNum         int       `gorm:"column:room_num" json:"room_num"`                   // 房间号
	DisplayContent  *string   `gorm:"column:display_content" json:"display_content"`     // 房屋展示内容json
	Floor           string    `gorm:"column:floor" json:"floor"`                         // 房屋所在楼层
	MaxFloor        string    `gorm:"column:max_floor" json:"max_floor"`                 // 房屋最大楼层
	HasElevator     int8      `gorm:"column:has_elevator" json:"has_elevator"`           // 是否有电梯（0没有、1有)
	BuildYear       string    `gorm:"column:build_year" json:"build_year"`               // 建成年份
	Direction       string    `gorm:"column:direction" json:"direction"`                 // 房屋朝向
	Lighting        int       `gorm:"column:lighting" json:"lighting"`                   // 房源采光情况
	NearTrafficJson string    `gorm:"column:near_traffic_json" json:"near_traffic_json"` // 附近交通信息
	CertificateNo   string    `gorm:"column:certificate_no" json:"certificate_no"`       // 房产证号
	LocationInfo    string    `gorm:"column:location_info" json:"location_info"`         // 房源地理位置信息
	JsonExtend      *string   `gorm:"column:json_extend"`                                // 扩展字段
	CreatedAt       time.Time `gorm:"column:created_at"`                                 // 创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at"`                                 // 更新时间
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

type HomeHouseDataResponse struct {
	// WholeHouseList 整租房源列表
	WholeHouseList []HouseListItem `json:"whole_house_list"`
	// ShareHouseList 合租房源列表
	ShareHouseList []HouseListItem `json:"share_house_list"`
}

// HouseListItem 房源列表项信息
type HouseListItem struct {
	HouseID     int64  `json:"house_id"`   // HouseID 房源编号
	Title       string `json:"title"`      // Title 房源标题
	IndexImg    string `json:"index_img"`  // IndexImg 房源封面图
	MonthlyRent int    `json:"rent_money"` // MonthlyRent 月租金

	HouseStateField HouseState `json:"state"`      // State 房源状态
	LeaseTypeField  RentType   `json:"rent_type"`  // RentType 租赁类型
	HouseTypeField  HouseType  `json:"house_type"` // HouseType 房源类型
	RentStateField  RentState  `json:"rent_state"` // RentState 出租状态

	City          string `json:"city"`            // City 所在城市
	District      string `json:"district"`        // District 所在区域
	Address       string `json:"address"`         // Address 详细地址
	Area          int    `json:"area"`            // Area 房源面积
	BedRoomNum    int    `json:"bed_room_num"`    // BedRoomNum 卧室数量
	LivingRoomNum int    `json:"living_room_num"` // LivingRoomNum 客厅数量
	ToiletNum     int    `json:"toilet_num"`      // ToiletNum 卫生间数量
	KitchenNum    int    `json:"kitchen_num"`     // KitchenNum 厨房数量
}

// HouseState 房源状态
type HouseState string

const (
	Up       HouseState = "up"       // up 房源状态: 已上架
	Down     HouseState = "down"     // down 房源状态: 已下架
	Auditing HouseState = "auditing" // auditing 房源状态: 审核中
	Deleted  HouseState = "deleted"  // deleted 房源状态: 已删除
)

type RentType string

const (
	Whole RentType = "whole" // whole 整租
	Share RentType = "share" // share 合租
)

func (t RentType) String() string {
	return string(t)
}

// HouseType 房屋类型
type HouseType string

const (
	Department  HouseType = "department"  // department 公寓
	Community   HouseType = "community"   // community 小区
	Residential HouseType = "residential" // residential 普通住宅
)

func (t HouseType) String() string {
	return string(t)
}

// RentState 出租状态
type RentState string

const (
	Rent    RentState = "rent"     // rent 已出租
	NotRent RentState = "not_rent" // not_rent 未出租
	Ordered RentState = "ordered"  // ordered 已预定
)

type HouseListRequest struct {
	QueryParams struct {
		RentMoneyRange []string `json:"rent_money_range"` // 月租金范围
		AreaRange      []string `json:"area_range"`       // 房屋面积范围
		Address        string   `json:"address"`          // 房源地址
		City           string   `json:"city"`             // 所在城市
		District       string   `json:"district"`         // 所在区县
		RentType       string   `json:"rent_type"`        // 租赁类型
		HouseOwner     int      `json:"house_owner"`      //	房屋拥有者
	} `json:"query_params"` // QueryParams 房源列表查询参数
	Offset int `json:"offset"` // Offset 分页偏移量
	Limit  int `json:"limit"`  // Limit 每页显示数量
}

type PublishHouseRequest struct {
	Title           string `json:"title"`            // Title 房源标题
	Address         string `json:"address"`          // Address 详细地址
	RentMoney       int    `json:"rent_money"`       // RentMoney 月租金
	City            string `json:"city"`             // City 所在城市
	District        string `json:"district"`         // District 所在区域
	BedroomNum      int    `json:"bedroom_num"`      // 卧室数量
	LivingRoomNum   int    `json:"living_room_num"`  // 客厅数量
	WaterRent       int    `json:"water_rent"`       // 水费(元/月)
	ElectricityRent int    `json:"electricity_rent"` // 电费(元/月)
	StrateRee       int    `json:"strate_fee"`       // 物业费(元/月)
	DepositRatio    int    `json:"deposit_ratio"`    // 租赁费用的押金倍数 (押几付几)
	PayRatio        int    `json:"pay_ratio"`        // 租赁费用的付款倍数 (押几付几)
	BargainMoney    int    `json:"bargain_money"`    // 定金
	CertificateNo   string `json:"certificate_no"`   // 房产证号

	RentTimeUnitField RentTimeUnit `json:"rent_time_unit"` // 租赁时间单位
	RentTypeField     RentType     `json:"rent_type"`      // 租赁类型
	HouseTypeField    HouseType    `json:"house_type"`     // 房源类型

	IndexImg        string `json:"index_img"`         // IndexImg 房源封面图
	HouseOwner      int    `json:"house_owner"`       // 房屋拥有者
	HouseDesc       string `json:"house_desc"`        // 房源描述
	Area            int    `json:"area"`              // 房源面积
	RoomNum         int    `json:"room_num"`          // 房间号
	ToiletNum       int    `json:"toilet_num"`        // 卫生间数量
	KitchenNum      int    `json:"kitchen_num"`       // 厨房数量
	Floor           string `json:"floor"`             // 楼层
	MaxFloor        string `json:"max_floor"`         // 最高楼层
	BuildYear       string `json:"build_year"`        // 建成年份
	NearTrafficJson string `json:"near_traffic_json"` // 附近交通

	HasElevatorField  HouseElevator           `json:"has_elevator"`        // 是否有电梯
	Direction         HouseDirection          `json:"direction"`           // 朝向
	LocationInfo      string                  `json:"location_info"`       // 房源地理位置信息
	DisplayContent    HouseDisplayContentItem `json:"display_content"`     // 房源展示内容
	HouseFacilityList []HouseFacilityListItem `json:"house_facility_list"` // 房源设施列表
	HouseContactInfo  HouseContactDataItem    `json:"house_contact_info"`  // 房源联系人信息
}

type RentTimeUnit string

const (
	Day      RentTimeUnit = "day"       // day 按天
	Month    RentTimeUnit = "month"     // month 按月
	Quarter  RentTimeUnit = "quarter"   // quarter 按季度
	HalfYear RentTimeUnit = "half_year" // half_year 按半年
	Year     RentTimeUnit = "year"      // year 按年
)

func (t RentTimeUnit) String() string {
	return string(t)
}

type HouseElevator int8

const (
	No  HouseElevator = 0 // no 无电梯
	Yes HouseElevator = 1 // yes 有电梯
)

type HouseDirection string

const (
	East  HouseDirection = "east"  // east 东
	West  HouseDirection = "west"  // west 西
	South HouseDirection = "south" // south 南
	North HouseDirection = "north" // north 北
)

func (h HouseDirection) String() string {
	return string(h)
}

type HouseLocation string

const (
	Nl HouseLocation = "nl" // nl 北纬
	Sl HouseLocation = "sl" // sl 南纬
)

func (h HouseLocation) Json() *string {
	location := string(h)
	parts := strings.Split(location, ",")
	logger.Debugw("HouseLocation.Json", "location", location, "parts", parts)
	str := fmt.Sprintf(`{"nl":%s,"sl":%s}`, parts[0], parts[1])
	return &str
}

type HouseDisplayContentItem string

const (
	Images HouseDisplayContentItem = "images" // images 房源图片
	Videos HouseDisplayContentItem = "videos" // video 房源视频
)

type HouseFacilityListItem struct {
	FacilityId int    `json:"facility_id"` // FacilityId 设施ID
	Name       string `json:"name"`        // Name 设施名称
	Icon       string `json:"icon"`        // icon 设施图标
}

type HouseContactDataItem struct {
	Mobile   string `json:"mobile"`    // Mobile 联系人手机号
	UserID   int    `json:"user_id"`   // UserID 联系人ID
	RealName string `json:"real_name"` // RealName 联系人姓名
	Email    string `json:"email"`     // Email 联系人邮箱
}

type HouseDetailDataItem struct {
	HouseListItem
	HouseOwner        int                     `json:"house_owner"`
	ContactID         int                     `json:"contact_id"`
	BargainMoney      int                     `json:"bargain_money"`
	WaterRent         int                     `json:"water_rent"`
	ElectricityRent   int                     `json:"electricity_rent"`
	StrataFee         int                     `json:"strata_fee"`
	DepositRatio      int                     `json:"deposit_ratio"`
	PayRatio          int                     `json:"pay_ratio"`
	HouseDesc         string                  `json:"house_desc"`
	Area              int                     `json:"area"`
	RoomNum           int                     `json:"room_num"`
	ToiletNum         int                     `json:"toilet_num"`
	Floor             int                     `json:"floor"`
	MaxFloor          int                     `json:"max_floor"`
	BuildYear         string                  `json:"build_year"`
	CertificateNo     string                  `json:"certificate_no"`
	NearTrafficJSON   map[string]interface{}  `json:"near_traffic_json"`
	RentTimeUnit      string                  `json:"rent_time_unit"`
	HasElevator       int8                    `json:"has_elevator"`
	DisplayContent    map[string]interface{}  `json:"display_content"`
	Direction         string                  `json:"direction"`
	LocationInfo      Location                `json:"location_info"`
	HouseFacilityList []HouseFacilityListItem `json:"house_facility_list"`
	HouseContactInfo  HouseContactDataItem    `json:"house_contact_info"`
}

type Location struct {
	Nl string `json:"nl"`
	Sl string `json:"sl"`
}
