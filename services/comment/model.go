package comment

// 房源评论结构体
type HouseComment struct {
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
