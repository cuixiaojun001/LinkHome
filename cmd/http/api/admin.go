package api

import (
	"github.com/cuixiaojun001/LinkHome/common/response"
	"github.com/cuixiaojun001/LinkHome/services/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserList(c *gin.Context) {
	var req admin.UserListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
		return
	}

	mgr := admin.GetAdminManager()
	resp, err := mgr.GetUserList(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.InternalServerError(err))
		return
	}

	c.JSON(http.StatusOK, response.Success(resp))
}

func OrderList(c *gin.Context) {
	var req admin.OrderListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
		return
	}

	mgr := admin.GetAdminManager()
	resp, err := mgr.GetOrderList(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.InternalServerError(err))
		return
	}

	c.JSON(http.StatusOK, response.Success(resp))
}
