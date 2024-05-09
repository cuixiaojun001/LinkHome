package house

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/cuixiaojun001/linkhome/common/logger"
	"github.com/cuixiaojun001/linkhome/modules/house/dao"
	"github.com/cuixiaojun001/linkhome/modules/house/model"
)

type IHouseService interface {
	// PublishHouse 发布房源
	PublishHouse(req PublishHouseRequest) (error, error)
	// GetHouseDetail 获取房源详情
	GetHouseDetail(id int) (*HouseDetail, error)
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

func HouseListInfo(req HouseListRequest) (*HouseListDataItem, error) {
	filter := req.GenQuery()
	houseList, err := dao.GetHouseList(filter)
	if err != nil {
		logger.Errorw("GetHouseList failed", "err", err)
		return nil, err
	}
	byteHouseList, _ := json.Marshal(houseList)
	var DataItem []HouseSummary
	_ = json.Unmarshal(byteHouseList, &DataItem)

	return &HouseListDataItem{
		DataList: DataItem,
		Total:    len(houseList),
	}, nil
}
