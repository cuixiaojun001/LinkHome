package dao

import (
	"github.com/cuixiaojun001/LinkHome/common/mysql"
	"github.com/cuixiaojun001/LinkHome/library/orm"
	"github.com/cuixiaojun001/LinkHome/modules/order/model"
	"gorm.io/gorm"
)

func CreateOrder(order *model.OrderModel) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Create(order).Error
}

func UpdateOrder(order *model.OrderModel) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Save(order).Error
}

func GetOrder(orderID int) (*model.OrderModel, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	order := &model.OrderModel{}
	err := db.Where("id = ?", orderID).First(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func GetUserOrderList(userID int) (orders []model.OrderModel, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	err = db.Where("tenant_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return
}

func CheckRecordExists(query orm.IQuery) (exist bool, err error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	user := &model.OrderModel{} // 初始化指针

	err = orm.SetQuery(db, query).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil // 没有找到记录
		}
		return false, err // 其他错误
	}
	return true, nil // 找到了记录
}
