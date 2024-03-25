package dao

import (
	"github.com/cuixiaojun001/linkhome/common/mysql"
	"github.com/cuixiaojun001/linkhome/modules/house/model"
	"strconv"
)

func AddHouse(house *model.HouseInfo) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Create(house).Error
}

func AddHouseDetail(houseDetail *model.HouseDetailInfo) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Create(houseDetail).Error
}

// GetRecentHouse 根据租赁类型和所在城市获取最新未出租房源，默认返回6条
func GetRecentHouse(rentType model.RentType, city string, limit int) (houses []model.HouseInfo, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	if city == "" {
		db = db.Where("rent_type = ? AND rent_state = ?", rentType, model.NotRent).Order("published_at desc").Limit(limit).Find(&houses)
	} else {
		db = db.Where("rent_type = ? AND city = ? AND rent_state = ?", rentType, city, model.NotRent).Order("published_at desc").Limit(limit).Find(&houses)
	}
	if db.Error != nil {
		return nil, db.Error
	}
	return
}

func GetHouseInfoByID(id int) (house model.HouseInfo, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	err = db.Where("id = ?", id).First(&house).Error
	return
}

func GetHouseDetailByID(id int) (houseDetail model.HouseDetailInfo, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	err = db.Where("house_id = ?", id).First(&houseDetail).Error
	return
}

func GetHouseList(request model.HouseListRequest) (results []model.HouseInfo, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	if len(request.QueryParams.RentMoneyRange) == 2 {
		min, _ := strconv.Atoi(request.QueryParams.RentMoneyRange[0])
		max, _ := strconv.Atoi(request.QueryParams.RentMoneyRange[1])
		db = db.Where("rent_money BETWEEN ? AND ?", min, max)
	}

	if len(request.QueryParams.AreaRange) == 2 {
		min, _ := strconv.Atoi(request.QueryParams.AreaRange[0])
		max, _ := strconv.Atoi(request.QueryParams.AreaRange[1])
		db = db.Where("area BETWEEN ? AND ?", min, max)
	}

	if request.QueryParams.Address != "" {
		db = db.Where("address LIKE ?", "%"+request.QueryParams.Address+"%")
	}

	if request.QueryParams.City != "" {
		db = db.Where("city LIKE ?", "%"+request.QueryParams.City+"%")
	}

	if request.QueryParams.District != "" {
		db = db.Where("district = ?", request.QueryParams.District)
	}

	if request.QueryParams.RentType != "" {
		db = db.Where("rent_type = ?", request.QueryParams.RentType)
	}

	if request.QueryParams.HouseOwner != 0 {
		db = db.Where("house_owner = ?", request.QueryParams.HouseOwner)
	}

	db = db.Limit(request.Limit).Offset(request.Offset)

	// Execute the query and get the result
	err = db.Find(&results).Error
	return results, err
}

func GetFacilityInfoByIDs(ids []int) ([]model.Facility, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	var facilities []model.Facility
	if err := db.Where("id IN (?)", ids).Find(&facilities).Error; err != nil {
		return nil, err
	}
	return facilities, nil
}

func GetFacilityByHouseID(houseId int) ([]model.Facility, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	var mappings []model.HouseFacilityMapping
	if err := db.Where("house_id = ?", houseId).Find(&mappings).Error; err != nil {
		return nil, err
	}

	var facilityIds []int
	for _, mapping := range mappings {
		facilityIds = append(facilityIds, mapping.FacilityID)
	}

	return GetFacilityInfoByIDs(facilityIds)
}
