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

func GetHouseInfo(houseID int) (*model.HouseInfo, error) {
	db := mysql.GetGormDB(mysql.MasterDB)
	var house model.HouseInfo
	err := db.Where("id = ?", houseID).First(&house).Error
	if err != nil {
		return nil, err
	}
	return &house, nil
}

func CreateHouseInfo(house *model.HouseInfo) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Create(house).Error
}

func UpdateHouseInfo(house *model.HouseInfo) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Save(house).Error
}

func CreateHouseDetail(houseDetail *model.HouseDetailInfo) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Create(houseDetail).Error
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
	query.SetOrder("rent_money", false)
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

// GetUserViewedHouses 获取当前用户在user_views表中浏览过的房源id列表
func GetUserViewedHouses(userID int) ([]int, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)

	var userViews []model.UserView
	if err := db.Where("user_id = ?", userID).Find(&userViews).Error; err != nil {
		return nil, err
	}

	var houseIDs []int
	for _, userView := range userViews {
		houseIDs = append(houseIDs, userView.HouseID)
	}
	return houseIDs, nil
}

// GetSimilarUsers 获取与当前用户有相似浏览房源的用户
func GetSimilarUsers(userID int, houseIDs []int) ([]int, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)

	var similarUsers []int
	if err := db.Table("user_views").Where("house_id IN (?) AND user_id != ?", houseIDs, userID).
		Group("user_id").
		Order("COUNT(*) DESC").
		Limit(10).
		Pluck("user_id", &similarUsers).Error; err != nil {
		return nil, err
	}

	return similarUsers, nil
}

// GetRecommendedHouses 获取推荐的房源
// 找出被相似用户浏览过但当前用户未浏览过的房源，并按照被浏览的次数进行排序，最后返回前10个房源。
func GetRecommendedHouses(userID int, similarUsers []int) ([]int, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)

	var userViews []model.UserView
	if err := db.Table("user_views").
		Select("house_id").
		Where("user_id IN (?) AND house_id NOT IN (SELECT house_id FROM user_views WHERE user_id = ?)", similarUsers, userID).
		Group("house_id").
		Order("COUNT(*) DESC").
		Limit(10).
		Find(&userViews).Error; err != nil {
		return nil, err
	}

	var recommendations []int
	for _, userView := range userViews {
		recommendations = append(recommendations, userView.HouseID)
	}
	return recommendations, nil
}
