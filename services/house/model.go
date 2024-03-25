package house

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

	HouseStateField HouseState `json:"state"`      // HouseStateField 房源状态
	LeaseTypeField  RentType   `json:"rent_type"`  // LeaseTypeField 租赁类型
	HouseTypeField  HouseType  `json:"house_type"` // HouseTypeField 房源类型
	RentStateField  RentState  `json:"rent_state"` // RentStateField 出租状态

	City          string `json:"city"`            // City 所在城市
	District      string `json:"district"`        // District 所在区域
	Address       string `json:"address"`         // Address 详细地址
	Area          int    `json:"area"`            // Area 房源面积
	BedRoomNum    int    `json:"bed_room_num"`    // BedRoomNum 卧室数量
	LivingRoomNum int    `json:"living_room_num"` // LivingRoomNum 客厅数量
	ToiletNum     int    `json:"toilet_num"`      // ToiletNum 卫生间数量
	KitchenNum    int    `json:"kitchen_num"`     // KitchenNum 厨房数量
}

type HouseState string

const (
	up       HouseState = "up"       // up 房源状态: 已上架
	down     HouseState = "down"     // down 房源状态: 已下架
	auditing HouseState = "auditing" // auditing 房源状态: 审核中
	deleted  HouseState = "deleted"  // deleted 房源状态: 已删除
)

type RentType string

const (
	whole RentType = "whole" // whole 整租
	share RentType = "share" // share 合租
)

type HouseType string

const (
	department  HouseType = "department"  // department 公寓
	community   HouseType = "community"   // community 小区
	residential HouseType = "residential" // residential 普通住宅
)

type RentState string

const (
	rent     RentState = "rent"     // rent 已出租
	not_rent RentState = "not_rent" // not_rent 未出租
	ordered  RentState = "ordered"  // ordered 已预定
)

// HouseListDataItem 房源列表数据
type HouseListDataItem struct {
	Total      int             `json:"total"`       // Total 数据总数量
	DataList   []HouseListItem `json:"data_list"`   // DataList 房源列表数据
	HasMore    bool            `json:"has_more"`    // HasMore 是否有下一页
	NextOffset int             `json:"next_offset"` // NextOffset offset下次起步
}
