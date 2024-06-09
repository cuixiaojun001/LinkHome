package dao

import (
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/common/mysql"
	"github.com/cuixiaojun001/LinkHome/library/orm"
	"github.com/cuixiaojun001/LinkHome/modules/order/model"
	"gorm.io/gorm"
	"log"
)

func CreateOrder(order *model.OrderModel) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Create(order).Error
}

func UpdateOrder(order *model.OrderModel) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Save(order).Error
}

func DeleteOrder(orderID int) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Where("id = ?", orderID).Delete(&model.OrderModel{}).Error
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

// 更新订单主键id，用于支付宝无法对同一订单多次支付
func UpdateOrderID(order *model.OrderModel) (*model.OrderModel, error) {
	db := mysql.GetGormDB(mysql.MasterDB)
	tx := db.Begin()
	// 确保事务在出错时回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("Recovered in UpdateOrderID:", r)
		}
	}()
	// 新订单模型是order去掉id字段值
	newOrder := *order
	newOrder.ID = 0
	if err := DeleteOrder(order.ID); err != nil {
		tx.Rollback()
		logger.Errorw("UpdateOrderID", "DeleteOrder", err)
		return nil, err
	}
	if err := CreateOrder(&newOrder); err != nil {
		tx.Rollback()
		logger.Errorw("UpdateOrderID", "CreateOrder", err)
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		logger.Errorw("UpdateOrderID", "Commit", err)
		return nil, err
	}
	logger.Debugw("UpdateOrderID", "order_id", order.ID, "new_order_id", newOrder.ID)
	return &newOrder, nil
}
