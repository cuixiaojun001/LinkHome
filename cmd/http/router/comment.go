package router

import (
	"github.com/cuixiaojun001/LinkHome/cmd/http/api"
	"github.com/gin-gonic/gin"
)

// RegisterCommentAPI 评论模块
func RegisterCommentAPI(engine *gin.Engine) {
	g := engine.Group("/api/v1/comment")
	{
		g.POST("/publish", api.PublishComment)
		g.POST("/reply/publish", api.PublishReplyComment)
	}
}
