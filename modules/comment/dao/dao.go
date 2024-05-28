package dao

import (
	"github.com/cuixiaojun001/LinkHome/common/mysql"
	"github.com/cuixiaojun001/LinkHome/modules/comment/model"
)

// GetHouseCommentsByHouseID 根据房源id获取所有房源主评论
func GetHouseCommentsByHouseID(houseID int) ([]model.HouseComments, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	var comments []model.HouseComments
	err := db.Where("house_id = ?", houseID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// 根据主评论id获取所有追评
func GetHouseCommentRepliesByCommentID(commentID int) ([]model.HouseCommentReplies, error) {
	db := mysql.GetGormDB(mysql.SlaveDB)
	var replies []model.HouseCommentReplies
	err := db.Where("comment_id = ?", commentID).Find(&replies).Error
	if err != nil {
		return nil, err
	}
	return replies, nil
}
