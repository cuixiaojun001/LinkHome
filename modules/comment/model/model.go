package model

import "time"

// HouseComments 房源评论表
type HouseComments struct {
	ID         int       `json:"id" gorm:"id"`                   // 评论ID
	UserId     int       `json:"user_id" gorm:"user_id"`         // 评论者ID
	HouseId    int       `json:"house_id" gorm:"house_id"`       // 房源ID
	Comment    string    `json:"comment" gorm:"comment"`         // 评论内容
	Time       time.Time `json:"time" gorm:"time"`               // 评论时间
	CommentNum int       `json:"comment_num" gorm:"comment_num"` // 追评数量
	Like       int       `json:"like" gorm:"like"`               // 点赞数量
	CreatedAt  time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`
}

// TableName 表名称
func (*HouseComments) TableName() string {
	return "house_comments"
}

// HouseCommentReplies 房源追评表
type HouseCommentReplies struct {
	ID         int       `gorm:"id"`           // 追评ID
	CommentId  int       `gorm:"comment_id"`   // 主评论ID
	FromUserId int       `gorm:"from_user_id"` // 追评者ID
	ToUserId   int       `gorm:"to_user_id"`   // 被追评者ID
	Comment    string    `gorm:"comment"`      // 追评内容
	Time       time.Time `gorm:"time"`         // 评论时间
	Like       int       `gorm:"like"`         // 点赞数量
	CreatedAt  time.Time `gorm:"created_at"`
	UpdatedAt  time.Time `gorm:"updated_at"`
}

// TableName 表名称
func (*HouseCommentReplies) TableName() string {
	return "house_comment_replies"
}
