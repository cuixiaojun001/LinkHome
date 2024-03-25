package dao

import (
	"github.com/cuixiaojun001/linkhome/common/mysql"
	"github.com/cuixiaojun001/linkhome/modules/common/model"
)

func GetAllProvince() ([]model.Province, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	var provinceList []model.Province
	err := db.Where("parent_id IS NULL").Find(&provinceList).Error
	if err != nil {
		return nil, err
	}
	return provinceList, nil
}

func GetCityByID(id int) ([]model.City, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	var areaList []model.City
	err := db.Where("parent_id = ?", id).Find(&areaList).Error
	if err != nil {
		return nil, err
	}
	return areaList, nil
}

func GetDistrictByID(id int) ([]model.District, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	var areaList []model.District
	err := db.Where("parent_id = ?", id).Find(&areaList).Error
	if err != nil {
		return nil, err
	}
	return areaList, nil
}
