package house

import (
	"encoding/json"
	"fmt"
	"github.com/cuixiaojun001/linkhome/library/orm"
	"strconv"
	"strings"
)

type PublishHouseRequest struct {
	Title           string  `json:"title"`            // Title 房源标题
	Address         string  `json:"address"`          // Address 详细地址
	RentMoney       int     `json:"rent_money"`       // RentMoney 月租金
	City            string  `json:"city"`             // City 所在城市
	District        string  `json:"district"`         // District 所在区域
	BedroomNum      int     `json:"bedroom_num"`      // 卧室数量
	LivingRoomNum   int     `json:"living_room_num"`  // 客厅数量
	WaterRent       float64 `json:"water_rent"`       // 水费(元/月)
	ElectricityRent float64 `json:"electricity_rent"` // 电费(元/月)
	StrateRee       int     `json:"strate_fee"`       // 物业费(元/月)
	DepositRatio    int     `json:"deposit_ratio"`    // 租赁费用的押金倍数 (押几付几)
	PayRatio        int     `json:"pay_ratio"`        // 租赁费用的付款倍数 (押几付几)
	BargainMoney    float64 `json:"bargain_money"`    // 定金

	RentTimeUnitField string `json:"rent_time_unit"` // 租赁时间单位
	RentTypeField     string `json:"rent_type"`      // 租赁类型
	HouseTypeField    string `json:"house_type"`     // 房源类型

	IndexImg        string `json:"index_img"`         // IndexImg 房源封面图
	HouseOwner      int    `json:"house_owner"`       // 房屋拥有者
	HouseDesc       string `json:"house_desc"`        // 房源描述
	Area            int    `json:"area"`              // 房源面积
	RoomNum         int    `json:"room_num"`          // 房间号
	ToiletNum       int    `json:"toilet_num"`        // 卫生间数量
	KitchenNum      int    `json:"kitchen_num"`       // 厨房数量
	Floor           int    `json:"floor"`             // 楼层
	MaxFloor        int    `json:"max_floor"`         // 最高楼层
	BuildYear       string `json:"build_year"`        // 建成年份
	NearTrafficJson string `json:"near_traffic_json"` // 附近交通
	CertificateNo   string `json:"certificate_no"`    // 房产证号

	HasElevatorField  int                     `json:"has_elevator"`        // 是否有电梯, 0:无, 1:有
	Direction         string                  `json:"direction"`           // 朝向
	LocationInfo      HouseLocation           `json:"location_info"`       // 房源地理位置信息
	DisplayContent    string                  `json:"display_content"`     // 房源展示内容
	HouseFacilityList []HouseFacilityListItem `json:"house_facility_list"` // 房源设施列表
	HouseContactInfo  HouseContactDataItem    `json:"house_contact_info"`  // 房源联系人信息
}

// HouseFacilityListItem 房源设施列表
type HouseFacilityListItem struct {
	FacilityId int    `json:"facility_id"` // FacilityId 设施ID
	Name       string `json:"name"`        // Name 设施名称
	Icon       string `json:"icon"`        // icon 设施图标
}

// HouseContactDataItem 房源联系人信息
type HouseContactDataItem struct {
	Mobile   string `json:"mobile"`    // Mobile 联系人手机号
	UserID   int    `json:"user_id"`   // UserID 联系人ID
	RealName string `json:"real_name"` // RealName 联系人姓名
	Email    string `json:"email"`     // Email 联系人邮箱
}

type HouseLocation string

func (h HouseLocation) Json() json.RawMessage {
	location := string(h)
	parts := strings.Split(location, ",")
	if len(parts) < 2 {
		return nil
	}
	str := fmt.Sprintf(`{"nl":%s,"sl":%s}`, parts[0], parts[1])
	return json.RawMessage(str)
}

// HouseSummary 房源摘要信息
type HouseSummary struct {
	HouseID     int    `json:"house_id"`   // HouseID 房源编号
	Title       string `json:"title"`      // Title 房源标题
	IndexImg    string `json:"index_img"`  // IndexImg 房源封面图
	MonthlyRent int    `json:"rent_money"` // MonthlyRent 月租金

	HouseStateField string `json:"state"`      // State 房源状态
	LeaseTypeField  string `json:"rent_type"`  // RentType 租赁类型
	HouseTypeField  string `json:"house_type"` // HouseType 房源类型
	RentStateField  string `json:"rent_state"` // RentState 出租状态

	City          string `json:"city"`            // City 所在城市
	District      string `json:"district"`        // District 所在区域
	Address       string `json:"address"`         // Address 详细地址
	Area          int    `json:"area"`            // Area 房源面积
	BedRoomNum    int    `json:"bed_room_num"`    // BedRoomNum 卧室数量
	LivingRoomNum int    `json:"living_room_num"` // LivingRoomNum 客厅数量
	ToiletNum     int    `json:"toilet_num"`      // ToiletNum 卫生间数量
	KitchenNum    int    `json:"kitchen_num"`     // KitchenNum 厨房数量
}

// HouseDetail 房源详细信息
type HouseDetail struct {
	HouseSummary
	HouseOwner      int                    `json:"house_owner"`       // HouseOwner 房屋拥有者
	ContactID       int                    `json:"contact_id"`        // ContactID 房源联系人id
	BargainMoney    float64                `json:"bargain_money"`     // BargainMoney 房屋预定金 (单位/分，元/100)
	WaterRent       float64                `json:"water_rent"`        // WaterRent 水费 (单位/分，元/100)
	ElectricityRent float64                `json:"electricity_rent"`  // ElectricityRent 电费 (单位/分，元/100)
	StrataFee       int                    `json:"strata_fee"`        // StrataFee 管理费 (单位/分，元/100)
	DepositRatio    int                    `json:"deposit_ratio"`     // DepositRatio 租赁费用的押金倍数 (押几付几)
	PayRatio        int                    `json:"pay_ratio"`         // PayRatio 租赁费用的支付倍数 (押几付几)
	HouseDesc       string                 `json:"house_desc"`        // HouseDesc 房屋描述
	Area            int                    `json:"area"`              // Area 房间面积
	RoomNum         int                    `json:"room_num"`          // RoomNum 房间号
	ToiletNum       int                    `json:"toilet_num"`        // ToiletNum 卫生间数量
	Floor           int                    `json:"floor"`             // Floor 房屋所在楼层
	MaxFloor        int                    `json:"max_floor"`         // MaxFloor 房屋最大楼层
	BuildYear       string                 `json:"build_year"`        // BuildYear 建成年份
	CertificateNo   string                 `json:"certificate_no"`    // CertificateNo 房产证号
	NearTrafficJSON map[string]interface{} `json:"near_traffic_json"` // NearTrafficJSON 附近交通信息

	RentTimeUnit      string                  `json:"rent_time_unit"`      // RentTimeUnit 租赁时间单位，默认month（月结）
	HasElevator       int8                    `json:"has_elevator"`        // HasElevator 是否有电梯
	DisplayContent    map[string]interface{}  `json:"display_content"`     // DisplayContent 房屋展示内容
	Direction         string                  `json:"direction"`           // Direction 房屋朝向
	LocationInfo      Location                `json:"location_info"`       // LocationInfo 房源地址位置信息
	HouseFacilityList []HouseFacilityListItem `json:"house_facility_list"` // HouseFacilityList 房源设施数据
	HouseContactInfo  HouseContactDataItem    `json:"house_contact_info"`  // HouseContactInfo 房源联系人信息
}

type Location struct {
	Nl float64 `json:"nl"`
	Sl float64 `json:"sl"`
}

type HouseListRequest struct {
	QueryParams struct {
		RentMoneyRange []string `json:"rent_money_range"` // 月租金范围
		AreaRange      []string `json:"area_range"`       // 房屋面积范围
		Address        string   `json:"address"`          // 房源地址
		City           string   `json:"city"`             // 所在城市
		District       string   `json:"district"`         // 所在区县
		RentType       string   `json:"rent_type"`        // 租赁类型
		HouseOwner     int      `json:"house_owner"`      // 房屋拥有者
	} `json:"query_params"` // QueryParams 房源列表查询参数
	Offset int `json:"offset"` // Offset 分页偏移量
	Limit  int `json:"limit"`  // Limit 每页显示数量
}

func (r *HouseListRequest) GenQuery() orm.IQuery {
	search := orm.NewQuery()
	if len(r.QueryParams.RentMoneyRange) == 2 {
		min, _ := strconv.Atoi(r.QueryParams.RentMoneyRange[0])
		max, _ := strconv.Atoi(r.QueryParams.RentMoneyRange[1])
		search.Range("rent_money", min, max)
	}

	if len(r.QueryParams.AreaRange) == 2 {
		min, _ := strconv.Atoi(r.QueryParams.AreaRange[0])
		max, _ := strconv.Atoi(r.QueryParams.AreaRange[1])
		search.Range("area", min, max)
	}

	if r.QueryParams.Address != "" {
		search.FuzzyMatch("address", r.QueryParams.Address)
	}

	if r.QueryParams.City != "" {
		search.FuzzyMatch("city", r.QueryParams.City)
	}

	if r.QueryParams.District != "" {
		search.ExactMatch("district", r.QueryParams.District)
	}

	if r.QueryParams.RentType != "" {
		search.ExactMatch("rent_type", r.QueryParams.RentType)
	}

	if r.QueryParams.HouseOwner != 0 {
		search.ExactMatch("house_owner", r.QueryParams.HouseOwner)
	}
	search.SetPagination(r.Offset, r.Limit)
	return search
}

// HouseListDataItem 房源列表数据
type HouseListDataItem struct {
	Total      int            `json:"total"`       // Total 数据总数量
	DataList   []HouseSummary `json:"data_list"`   // DataList 房源列表数据
	HasMore    bool           `json:"has_more"`    // HasMore 是否有下一页
	NextOffset int            `json:"next_offset"` // NextOffset offset下次起步
}

type HouseFacilityListResponse struct {
	HouseFacilityList []HouseFacilityListItem `json:"house_facility_list"` // 房源设施列表
}
