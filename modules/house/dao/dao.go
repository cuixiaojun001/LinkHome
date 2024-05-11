package dao

import (
	"github.com/cuixiaojun001/LinkHome/common/mysql"
	"github.com/cuixiaojun001/LinkHome/library/orm"
	"github.com/cuixiaojun001/LinkHome/modules/house/model"
)

// GetRecentHouse 根据租赁类型和所在城市获取最新未出租房源，默认返回6条
func GetRecentHouse(rentType string, city string, limit int) (houses []model.HouseInfo, err error) {
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

func CreateHouseInfo(house *model.HouseInfo) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Create(house).Error
}

func CreateHouseDetail(houseDetail *model.HouseDetailInfo) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Create(houseDetail).Error
}

func GetHouseInfo(houseID int) (*model.HouseInfo, error) {
	db := mysql.GetGormDB(mysql.MasterDB)
	var house model.HouseInfo
	err := db.Where("id = ?", houseID).First(&house).Error
	if err != nil {
		return nil, err
	}
	return &house, nil
}

func GetHouseDetail(houseID int) (*model.HouseDetailInfo, error) {
	db := mysql.GetGormDB(mysql.MasterDB)
	var houseDetail model.HouseDetailInfo
	err := db.Where("id = ?", houseID).First(&houseDetail).Error
	if err != nil {
		return nil, err
	}
	return &houseDetail, nil
}

func FetchFacilitiesByHouseID(houseID int) ([]model.Facility, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	var mappings []model.HouseFacilityMapping
	if err := db.Where("house_id = ?", houseID).Find(&mappings).Error; err != nil {
		return nil, err
	}

	var facilityIds []int
	for _, mapping := range mappings {
		facilityIds = append(facilityIds, mapping.FacilityID)
	}

	return FetchFacilitiesInfoByIDs(facilityIds)
}

func FetchFacilitiesInfoByIDs(ids []int) ([]model.Facility, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	var facilities []model.Facility
	if err := db.Where("id IN (?)", ids).Find(&facilities).Error; err != nil {
		return nil, err
	}
	return facilities, nil
}

func GetHouseList(query orm.IQuery) (results []model.HouseInfo, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	err = orm.SetQuery(db, query).Find(&results).Error
	return results, err
}

func GetAllHouseFacility() ([]model.Facility, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	var facilities []model.Facility
	if err := db.Find(&facilities).Error; err != nil {
		return nil, err
	}
	return facilities, nil
}
