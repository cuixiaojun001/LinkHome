package dao

import (
	"github.com/cuixiaojun001/LinkHome/common/mysql"
	"github.com/cuixiaojun001/LinkHome/modules/payment/model"
)

func CreatePaymentTrade(paymentTrade *model.PaymentTrade) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Create(paymentTrade).Error
}
