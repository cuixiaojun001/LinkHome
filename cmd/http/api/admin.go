package api

import (
	"github.com/cuixiaojun001/LinkHome/services/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserList(c *gin.Context) {
	var req admin.UserListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mgr := admin.GetAdminManager()
	resp, err := mgr.GetUserList(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
