package dao

import (
	"github.com/cuixiaojun001/linkhome/common/mysql"
	"github.com/cuixiaojun001/linkhome/modules/house/model"
)

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
