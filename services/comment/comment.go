package comment

import (
	"context"
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/modules/comment/dao"
	"github.com/cuixiaojun001/LinkHome/modules/comment/model"
	userDao "github.com/cuixiaojun001/LinkHome/modules/user/dao"
	"sync"
	"time"
)

type ICommentService interface {
	// GetHouseCommentsByHouseID 根据房源id获取所有评论
	GetHouseCommentsByHouseID(ctx context.Context, houseID int) ([]HouseComment, error)
	// PublishComment 发布房源评论
	PublishComment(ctx context.Context, req *PublishCommentRequest) (*HouseComment, error)
	// PublishReplyComment 发布评论追评
	PublishReplyComment(ctx context.Context, req *PublishReplyCommentRequest) (*ReplyComment, error)
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
				CommentID:   reply.ID,
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
			CommentID:  main.ID,
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

func (s *CommentService) PublishComment(ctx context.Context, req *PublishCommentRequest) (*HouseComment, error) {
	comment := &model.HouseComments{
		UserId:     req.UserID,
		HouseId:    req.HouseID,
		Comment:    req.Comment,
		CommentNum: 0,
		Like:       0,
		Time:       time.Now(),
	}
	logger.Debugw("comment", "comment:", comment)
	if err := dao.CreateHouseComment(comment); err != nil {
		logger.Errorw("CreateHouseComment err ", "err:", err)
		return nil, err
	}
	return &HouseComment{
		CommentID: comment.ID,
	}, nil
}

func (s *CommentService) PublishReplyComment(ctx context.Context, req *PublishReplyCommentRequest) (*ReplyComment, error) {
	comment := &model.HouseCommentReplies{
		CommentId:  req.CommentID,
		FromUserId: req.FromUserID,
		ToUserId:   req.ToUserID,
		Comment:    req.Comment,
		Time:       time.Now(),
		Like:       0,
	}
	logger.Debugw("comment", "comment:", comment)
	if err := dao.CreateHouseCommentReply(comment); err != nil {
		logger.Errorw("CreateHouseComment err ", "err:", err)
		return nil, err
	}
	if err := dao.IncrementCommentNum(req.CommentID); err != nil {
		return nil, err
	}

	return &ReplyComment{
		CommentID: comment.ID,
	}, nil
}
