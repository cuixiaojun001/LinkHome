package comment

// 房源评论结构体
type HouseComment struct {
	CommentID  int            `json:"comment_id"` // CommentID 评论ID
	Name       string         `json:"name"`       // Name 评论者用户名
	ID         int            `json:"id"`         // ID 评论者Id
	HeadImg    string         `json:"headImg"`    // HeadImg 评论者头像
	Comment    string         `json:"comment"`    // Comment 评论内容
	Time       string         `json:"time"`       // Time 评论时间
	CommentNum int            `json:"commentNum"` // CommentNum 追评数量
	Like       int            `json:"like"`       // Like 点赞数量
	InputShow  bool           `json:"inputShow"`  // InputShow 是否显示输入框
	Reply      []ReplyComment `json:"reply"`      // Reply 追评
}

// ReplyComment 追评
type ReplyComment struct {
	CommentID   int    `json:"comment_id"`  // CommentID 评论ID
	From        string `json:"from"`        // From 追评者
	FromID      int    `json:"fromId"`      // FromID 追评者Id
	FromHeadImg string `json:"fromHeadImg"` // FromHeadImg 追评者头像
	To          string `json:"to"`          // To 被追评者
	ToID        int    `json:"toId"`        // ToID 被追评者Id
	Comment     string `json:"comment"`     // Comment 追评内容
	Time        string `json:"time"`        // Time 评论时间
	CommentNum  int    `json:"commentNum"`  // CommentNum 追评数量
	Like        int    `json:"like"`        // Like 点赞数量
	InputShow   bool   `json:"inputShow"`   // InputShow 是否显示输入框
}

type PublishCommentRequest struct {
	HouseID int    `json:"house_id"` // HouseID 房屋ID
	UserID  int    `json:"user_id"`  // UserID 评论用户ID
	Comment string `json:"comment"`  // Comment  评论内容
}

type PublishReplyCommentRequest struct {
	CommentID  int    `json:"comment_id"`   // CommentID 主评论ID
	FromUserID int    `json:"from_user_id"` // FromUserID 追评者ID
	ToUserID   int    `json:"to_user_id"`   // ToUserID 被追评者ID
	Comment    string `json:"comment"`      // Comment  评论内容
}
