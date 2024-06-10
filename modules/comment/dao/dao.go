package dao

import (
	"github.com/cuixiaojun001/LinkHome/common/mysql"
	"github.com/cuixiaojun001/LinkHome/library/orm"
	"github.com/cuixiaojun001/LinkHome/modules/comment/model"
)

func CreateHouseComment(comment *model.HouseComments) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Create(comment).Error
}

func CreateHouseCommentReply(reply *model.HouseCommentReplies) error {
	db := mysql.GetGormDB(mysql.MasterDB)
	return db.Create(reply).Error
}

func IncrementCommentNum(commentID int) error {
	db := mysql.GetGormDB(mysql.MasterDB)

	// 查找对应ID的房源评论
	var comment model.HouseComments
	filter := orm.NewQuery().ExactMatch("id", commentID)
	if err := orm.SetQuery(db, filter).Find(&comment).Error; err != nil {
		return err
	}

	// 更新评论数量
	comment.CommentNum++

	// 保存更新后的评论
	if err := db.Save(&comment).Error; err != nil {
		return err
	}

	return nil
}

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
