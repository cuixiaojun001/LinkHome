package comment

import (
	"context"
	"github.com/cuixiaojun001/LinkHome/modules/comment/dao"
	userDao "github.com/cuixiaojun001/LinkHome/modules/user/dao"
	"sync"
)

type ICommentService interface {
	// GetHouseCommentsByHouseID 根据房源id获取所有评论
	GetHouseCommentsByHouseID(ctx context.Context, houseID int) ([]HouseComment, error)
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

func (s *CommentService) GetHouseCommentsByHouseID(ctx context.Context, houseID int) ([]HouseComment, error) {
	comments := []HouseComment{}
	mainComments, err := dao.GetHouseCommentsByHouseID(houseID)
	if err != nil {
		return nil, err
	}

	for _, main := range mainComments {
		user, err := userDao.GetUserBasicInfo(main.UserId)
		if err != nil {
			return nil, err
		}
		replies, err := dao.GetHouseCommentRepliesByCommentID(main.ID)
		if err != nil {
			return nil, err
		}
		replyComment := []ReplyComment{}
		for _, reply := range replies {
			replyUser, err := userDao.GetUserBasicInfo(reply.FromUserId)
			if err != nil {
				continue
			}
			temp := ReplyComment{
				From:        replyUser.Username,
				FromID:      reply.FromUserId,
				FromHeadImg: "",
				To:          user.Username,
				ToID:        main.UserId,
				Comment:     reply.Comment,
				Time:        reply.Time.Format("2006-01-02"),
				CommentNum:  0, // FIXME 修改追评表
				Like:        reply.Like,
				InputShow:   false,
			}
			replyComment = append(replyComment, temp)
		}

		comment := HouseComment{
			Name:       user.Username,
			ID:         main.UserId,
			HeadImg:    "",
			Comment:    main.Comment,
			Time:       main.Time.Format("2006-01-02"),
			CommentNum: main.CommentNum,
			Like:       main.Like,
			InputShow:  false,
			Reply:      replyComment,
		}

		comments = append(comments, comment)
	}
	return comments, nil
}
