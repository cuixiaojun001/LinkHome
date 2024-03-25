package house

import (
	"encoding/json"
	"github.com/cuixiaojun001/linkhome/common/logger"
	"github.com/cuixiaojun001/linkhome/modules/house/dao"
	"github.com/cuixiaojun001/linkhome/modules/house/model"
	"github.com/cuixiaojun001/linkhome/third_party/qiniu"
	"strconv"
	"time"
)

// HomeHouseInfo 首页推荐房源信息
func HomeHouseInfo(city string) (*HomeHouseDataResponse, error) {
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
	var wholeHouseItem []HouseListItem
	_ = json.Unmarshal(wholeHouse, &wholeHouseItem)

	shareHouse, _ := json.Marshal(ShareHouseList)
	var shareHouseItem []HouseListItem
	_ = json.Unmarshal(shareHouse, &shareHouseItem)

	return &HomeHouseDataResponse{
		WholeHouseList: wholeHouseItem,
		ShareHouseList: shareHouseItem,
	}
}

// HouseListInfo 获取房源列表信息
func HouseListInfo(req model.HouseListRequest) (*HouseListDataItem, error) {
	houseList, err := dao.GetHouseList(req)
	if err != nil {
		logger.Errorw("GetHouseList failed", "err", err)
		return nil, err
	}
	byteHouseList, _ := json.Marshal(houseList)
	var DataItem []HouseListItem
	_ = json.Unmarshal(byteHouseList, &DataItem)

	return &HouseListDataItem{
		DataList: DataItem,
		Total:    len(houseList),
	}, nil
}

func PublishHouse(req model.PublishHouseRequest) (*HouseListDataItem, error) {
	logger.Debugw("CertificateNo", "cert", req.CertificateNo)
	houseInfo, houseDetail, err := publishHouse(req)
	if err != nil {
		return nil, err
	}

	logger.Debugw("publishHouse success", "houseInfo", houseInfo, "houseDetail", houseDetail)

	return nil, nil
}

func publishHouse(req model.PublishHouseRequest) (*model.HouseInfo, *model.HouseDetailInfo, error) {
	house := &model.HouseInfo{
		RentType:        req.RentTypeField.String(),
		HouseType:       req.HouseTypeField.String(),
		Title:           req.Title,
		IndexImg:        req.IndexImg,
		HouseDesc:       req.HouseDesc,
		City:            req.City,
		District:        req.District,
		Address:         req.Address,
		RentState:       string(model.NotRent),
		RentMoney:       req.RentMoney,
		BargainMoney:    req.BargainMoney,
		RentTimeUnit:    req.RentTimeUnitField.String(),
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
		State:           string(model.Auditing),
		// JsonExtend: nil,
		HouseOwner: req.HouseOwner,
	}

	logger.Debugw("publishHouse", "house", house)

	if err := dao.AddHouse(house); err != nil {
		logger.Errorw("AddHouse failed", "err", err)
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
		Direction:   req.Direction.String(),
		// Lighting
		// NearTrafficJson
		CertificateNo: req.CertificateNo,
		LocationInfo:  req.LocationInfo,
		// JsonExtend: 默认值（空）
	}

	logger.Debugw("publishHouse", "houseDetail", houseDetail)
	if err := dao.AddHouseDetail(houseDetail); err != nil {
		logger.Errorw("AddHouseDetail failed", "err", err)
		return nil, nil, err
	}

	return house, houseDetail, nil
}

func GetHouse(id int) (*model.HouseDetailDataItem, error) {
	logger.Debugw("GetHouse", "id", id)
	house, err := dao.GetHouseInfoByID(id)
	if err != nil {
		logger.Errorw("GetHouse failed", "err", err)
		return nil, err
	}

	houseDetail, err := dao.GetHouseDetailByID(id)
	if err != nil {
		logger.Errorw("GetHouseDetail failed", "err", err)
		return nil, err
	}

	//  获取房源设施信息
	houseFacilityList, err := dao.GetFacilityByHouseID(id)
	if err != nil {
		logger.Errorw("GetFacilityByHouseID failed", "err", err)
		return nil, err
	}

	tmp, _ := json.Marshal(house)
	var HouseItem model.HouseListItem
	_ = json.Unmarshal(tmp, &HouseItem)

	tmp2, _ := json.Marshal(houseFacilityList)
	var facilityList []model.HouseFacilityListItem
	_ = json.Unmarshal(tmp2, &facilityList)

	floor, _ := strconv.Atoi(houseDetail.Floor)
	maxFloor, _ := strconv.Atoi(houseDetail.MaxFloor)

	item := &model.HouseDetailDataItem{
		HouseListItem:     HouseItem,
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
		Floor:             floor,
		MaxFloor:          maxFloor,
		BuildYear:         houseDetail.BuildYear,
		CertificateNo:     houseDetail.CertificateNo,
		RentTimeUnit:      house.RentTimeUnit,
		HasElevator:       houseDetail.HasElevator,
		// DisplayContent:    houseDetail.DisplayContent,
		Direction:    houseDetail.Direction,
		LocationInfo: houseDetail.LocationInfo,
		// HouseContactInfo
	}
	return item, nil
}

func convertToLocation(locationInfo *string) model.Location {
	logger.Debugw("convertToLocation", "locationInfo", locationInfo)
	loc := model.Location{}
	_ = json.Unmarshal([]byte(*locationInfo), &loc)
	logger.Debugw("convertToLocation", "loc", loc)
	return loc
}
