package api

import (
	"context"
	"github.com/cuixiaojun001/LinkHome/common/response"
	"github.com/cuixiaojun001/LinkHome/services/comment"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublishComment(c *gin.Context) {
	req := &comment.PublishCommentRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}

	mgr := comment.GetCommentManager()
	if resp, err := mgr.PublishComment(context.Background(), req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(resp))
	}
}

func PublishReplyComment(c *gin.Context) {
	req := &comment.PublishReplyCommentRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}

	mgr := comment.GetCommentManager()
	if resp, err := mgr.PublishReplyComment(context.Background(), req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(resp))
	}
}
