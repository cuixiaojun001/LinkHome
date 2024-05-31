package dao

import (
	"github.com/cuixiaojun001/LinkHome/common/mysql"
	"github.com/cuixiaojun001/LinkHome/library/orm"
	"github.com/cuixiaojun001/LinkHome/modules/user/model"
)

func GetUserProfile(filter orm.IQuery) (profiles []model.UserProfileInfo, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	err = orm.SetQuery(db, filter).First(&profiles).Error
	return profiles, err
}

func GetUserBasic(filter orm.IQuery) (profiles []model.UserBasicInfo, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	err = orm.SetQuery(db, filter).First(&profiles).Error
	return profiles, err
}
