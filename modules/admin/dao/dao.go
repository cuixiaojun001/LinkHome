package dao

import (
	"github.com/cuixiaojun001/LinkHome/common/mysql"
	"github.com/cuixiaojun001/LinkHome/library/orm"
	orderModel "github.com/cuixiaojun001/LinkHome/modules/order/model"
	userModel "github.com/cuixiaojun001/LinkHome/modules/user/model"
)

func GetUserProfile(filter orm.IQuery) (profiles []userModel.UserProfileInfo, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	err = orm.SetQuery(db, filter).Find(&profiles).Error
	return profiles, err
}

func GetUserBasic(filter orm.IQuery) (profiles []userModel.UserBasicInfo, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	err = orm.SetQuery(db, filter).Find(&profiles).Error
	return profiles, err
}

func GetOrder(filter orm.IQuery) (orders []orderModel.OrderModel, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	err = orm.SetQuery(db, filter).Find(&orders).Error
	return orders, err
}
