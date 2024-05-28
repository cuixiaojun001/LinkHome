package comment

import (
	"context"
	"github.com/cuixiaojun001/LinkHome/modules/comment/dao"
	"sync"
)

type ICommentService interface {
	// GetHouseCommentsByHouseID 根据房源id获取所有评论
	GetHouseCommentsByHouseID(ctx context.Context, houseID int) (*HouseComment, error)
}

type CommentService struct{}

var once sync.Once
var commentManager ICommentService

func GetCommentManager() ICommentService {
	once.Do(func() {
		commentManager = &CommentService{}
	})
	return commentManager
}

var _ ICommentService = (*CommentService)(nil)

func (s *CommentService) GetHouseCommentsByHouseID(ctx context.Context, houseID int) (*HouseComment, error) {
	mainComments, err := dao.GetHouseCommentsByHouseID(houseID)
	if err != nil {
		return nil, err
	}

	for _, main := range mainComments {
		_, err := dao.GetHouseCommentRepliesByCommentID(main.ID)
		if err != nil {
			return nil, err
		}
		// main.Replies = replies
	}
	return nil, nil
}
