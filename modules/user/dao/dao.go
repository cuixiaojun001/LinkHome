package dao

import (
	"github.com/cuixiaojun001/linkhome/common/logger"
	"github.com/cuixiaojun001/linkhome/common/mysql"
	"github.com/cuixiaojun001/linkhome/modules/user/model"
)

// AddUser 注册用户，添加一条记录
func AddUser(user *model.UserBasicInfo) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Create(user).Error
}

// CreateOrUpdateUserRentalDemand 创建或更新用户租房需求
func CreateOrUpdateUserRentalDemand(demand *model.UserRentalDemandInfo, modelID int) error {
	logger.Debugw("CreateOrUpdateUserRentalDemand", "demand", demand, "modelID", modelID)
	db := mysql.GetGormDB(mysql.MasterDB)
	if modelID == 0 {
		return db.Create(demand).Error
	} else {
		info := &model.UserRentalDemandInfo{}
		return db.Table(info.TableName()).Where("id = ?", modelID).Updates(demand).Error
	}
}

func IsUsernameExist(username string) bool {
	db := mysql.GetGormDB(mysql.SlaveDB)
	var user model.UserBasicInfo
	db.Where("username = ?", username).First(&user)
	return user.ID != 0
}

func IsMobileExist(mobile string) bool {
	db := mysql.GetGormDB(mysql.SlaveDB)
	var user model.UserBasicInfo
	db.Where("mobile = ?", mobile).First(&user)
	return user.ID != 0
}

// FilterFirst 条件筛选取第一个
func FilterFirst(params map[string]string) (user model.UserBasicInfo, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	for k, v := range params {
		db = db.Where(k+" = ?", v)
	}
	err = db.Take(&user).Error
	return user, err
}

func UpdateUserPwd(id, newPassword string) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Where("id = ?", id).Model(model.UserBasicInfo{}).Update("password", newPassword).Error
}

func GetUserBasicInfoByID(id string) (user model.UserBasicInfo, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	err = db.Where("id = ?", id).First(&user).Error
	return
}

func GetUserProfileByID(id string) (profile model.UserProfileInfo, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	err = db.Where("id = ?", id).First(&profile).Error
	return
}
