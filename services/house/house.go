package house

import (
	"encoding/json"
	"github.com/cuixiaojun001/linkhome/third_party/qiniu"
	"sync"
	"time"

	"github.com/cuixiaojun001/linkhome/common/logger"
	"github.com/cuixiaojun001/linkhome/modules/house/dao"
	"github.com/cuixiaojun001/linkhome/modules/house/model"
)

type IHouseService interface {
	// HomeHouseInfo 首页推荐房源信息
	HomeHouseInfo(city string) (*HomeHouseDataResponse, error)
	// PublishHouse 发布房源
	PublishHouse(req PublishHouseRequest) (error, error)
	// GetHouseDetail 获取房源详情
	GetHouseDetail(id int) (*HouseDetail, error)
	// HouseListInfo 获取房源列表
	HouseListInfo(req HouseListRequest) (*HouseListDataItem, error)
	// GetAllHouseFacility 获取所有房源设施
	GetAllHouseFacility() (*HouseFacilityListResponse, error)
}

type HouseService struct{}

var once sync.Once
var houseManager IHouseService

func GetHouseManager() IHouseService {
	once.Do(func() {
		houseManager = &HouseService{}
	})
	return houseManager
}

var _ IHouseService = (*HouseService)(nil)

func (s *HouseService) HomeHouseInfo(city string) (*HomeHouseDataResponse, error) {
	WholeHouseList, err := dao.GetRecentHouse(model.Whole, city, 6)
	if err != nil {
		return nil, err
	}
	SharedHouseList, err := dao.GetRecentHouse(model.Share, city, 6)
	if err != nil {
		return nil, err
	}

	// index_img
	for i := range WholeHouseList {
		WholeHouseList[i].IndexImg = qiniu.Client.MakePrivateURL(WholeHouseList[i].IndexImg)
	}

	for i := range SharedHouseList {
		SharedHouseList[i].IndexImg = qiniu.Client.MakePrivateURL(SharedHouseList[i].IndexImg)
	}

	result := newHomeHouseDataResponse(WholeHouseList, SharedHouseList)

	return result, nil
}

func newHomeHouseDataResponse(WholeHouseList, ShareHouseList []model.HouseInfo) *HomeHouseDataResponse {
	wholeHouse, _ := json.Marshal(WholeHouseList)
	var wholeHouseItem []HouseSummary
	_ = json.Unmarshal(wholeHouse, &wholeHouseItem)

	shareHouse, _ := json.Marshal(ShareHouseList)
	var shareHouseItem []HouseSummary
	_ = json.Unmarshal(shareHouse, &shareHouseItem)

	return &HomeHouseDataResponse{
		WholeHouseList: wholeHouseItem,
		ShareHouseList: shareHouseItem,
	}
}

func (s *HouseService) PublishHouse(req PublishHouseRequest) (error, error) {
	houseInfo, houseDetail, err := publishHouse(req)
	if err != nil {
		return nil, err
	}

	logger.Debugw("PublishHouse success", "houseInfo", houseInfo, "houseDetail", houseDetail)
	return nil, nil
}

func publishHouse(req PublishHouseRequest) (*model.HouseInfo, *model.HouseDetailInfo, error) {
	house := &model.HouseInfo{
		RentType:        req.RentTypeField,
		HouseType:       req.HouseTypeField,
		Title:           req.Title,
		IndexImg:        req.IndexImg,
		HouseDesc:       req.HouseDesc,
		City:            req.City,
		District:        req.District,
		Address:         req.Address,
		RentState:       model.NotRent,
		RentMoney:       req.RentMoney,
		BargainMoney:    req.BargainMoney,
		RentTimeUnit:    req.RentTimeUnitField,
		WaterRent:       req.WaterRent,
		ElectricityRent: req.ElectricityRent,
		StrataFee:       req.StrateRee,
		DepositRatio:    req.DepositRatio,
		PayRatio:        req.PayRatio,
		BedroomNum:      req.BedroomNum,
		LivingRoomNum:   req.LivingRoomNum,
		KitchenNum:      req.KitchenNum,
		ToiletNum:       req.ToiletNum,
		Area:            req.Area,
		PublishedAt:     time.Now(),
		State:           model.Auditing,
		// JsonExtend: nil,
		HouseOwner: req.HouseOwner,
	}

	if err := dao.CreateHouseInfo(house); err != nil {
		logger.Errorw("CreateHouseInfo failed", "err", err)
		return nil, nil, err
	}

	houseDetail := &model.HouseDetailInfo{
		HouseID:    house.ID,
		HouseOwner: req.HouseOwner,
		ContactId:  req.HouseContactInfo.UserID,
		Address:    req.Address,
		RoomNum:    req.RoomNum,
		// DisplayContent: req.DisplayContent,
		Floor:       req.Floor,
		MaxFloor:    req.MaxFloor,
		HasElevator: int8(req.HasElevatorField),
		BuildYear:   req.BuildYear,
		Direction:   req.Direction,
		// Lighting
		// NearTrafficJson
		CertificateNo: req.CertificateNo,
		LocationInfo:  req.LocationInfo.Json(),
		// JsonExtend: 默认值（空）
	}

	if err := dao.CreateHouseDetail(houseDetail); err != nil {
		logger.Errorw("CreateHouseDetail failed", "err", err)
		return nil, nil, err
	}

	return house, houseDetail, nil
}

func (s *HouseService) GetHouseDetail(id int) (*HouseDetail, error) {
	house, err := dao.GetHouseInfo(id)
	if err != nil {
		logger.Errorw("GetHouse failed", "err", err)
		return nil, err
	}

	houseDetail, err := dao.GetHouseDetail(id)
	if err != nil {
		logger.Errorw("GetHouseDetail failed", "err", err)
		return nil, err
	}

	//  获取房源设施信息
	houseFacilityList, err := dao.FetchFacilitiesByHouseID(id)
	if err != nil {
		logger.Errorw("GetFacilityByHouseID failed", "err", err)
		return nil, err
	}

	tmp, _ := json.Marshal(house)
	var summary HouseSummary
	_ = json.Unmarshal(tmp, &summary)
	summary.IndexImg = qiniu.Client.MakePrivateURL(summary.IndexImg)

	tmp2, _ := json.Marshal(houseFacilityList)
	var facilityList []HouseFacilityListItem
	_ = json.Unmarshal(tmp2, &facilityList)

	detail := &HouseDetail{
		HouseSummary:      summary,
		HouseFacilityList: facilityList,
		BargainMoney:      house.BargainMoney,
		WaterRent:         house.WaterRent,
		ElectricityRent:   house.ElectricityRent,
		StrataFee:         house.StrataFee,
		DepositRatio:      house.DepositRatio,
		PayRatio:          house.PayRatio,
		HouseDesc:         house.HouseDesc,
		Area:              house.Area,
		RoomNum:           houseDetail.RoomNum,
		ToiletNum:         house.ToiletNum,
		Floor:             houseDetail.Floor,
		MaxFloor:          houseDetail.MaxFloor,
		BuildYear:         houseDetail.BuildYear,
		CertificateNo:     houseDetail.CertificateNo,
		RentTimeUnit:      house.RentTimeUnit,
		HasElevator:       houseDetail.HasElevator,
		// DisplayContent:    houseDetail.DisplayContent,
		Direction:    houseDetail.Direction,
		LocationInfo: convertToLocation(houseDetail.LocationInfo),
		// HouseContactInfo
	}
	return detail, nil
}

func convertToLocation(rawJson json.RawMessage) Location {
	logger.Debugw("convertToLocation", "locationInfo", rawJson)
	var loc Location
	err := json.Unmarshal(rawJson, &loc)
	if err != nil {
		logger.Errorw("convertToLocation failed", "err", err)
		return Location{}
	}
	return loc
}

func (s *HouseService) HouseListInfo(req HouseListRequest) (*HouseListDataItem, error) {
	filter := req.GenQuery()
	houseList, err := dao.GetHouseList(filter)
	if err != nil {
		logger.Errorw("GetHouseList failed", "err", err)
		return nil, err
	}
	byteHouseList, _ := json.Marshal(houseList)
	var summary []HouseSummary
	_ = json.Unmarshal(byteHouseList, &summary)
	for i := range summary {
		summary[i].IndexImg = qiniu.Client.MakePrivateURL(summary[i].IndexImg)
	}

	return &HouseListDataItem{
		DataList: summary,
		Total:    len(houseList),
	}, nil
}

func (s *HouseService) GetAllHouseFacility() (*HouseFacilityListResponse, error) {
	facilityList, err := dao.GetAllHouseFacility()
	if err != nil {
		logger.Errorw("GetAllHouseFacility failed", "err", err)
		return nil, err
	}
	tmp, _ := json.Marshal(facilityList)
	var items []HouseFacilityListItem
	_ = json.Unmarshal(tmp, &items)
	for i := range items {
		items[i].Icon = qiniu.Client.MakePrivateURL(items[i].Icon)
	}

	return &HouseFacilityListResponse{HouseFacilityList: items}, nil
}