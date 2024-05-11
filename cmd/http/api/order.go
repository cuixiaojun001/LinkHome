package api

import (
	"context"
	"github.com/cuixiaojun001/linkhome/common/response"
	"github.com/cuixiaojun001/linkhome/services/order"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateOrder(c *gin.Context) {
	userID := c.Param("user_id")
	// 转为int
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	var req order.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	mgr := order.GetOrderManager()
	if err := mgr.CreateOrder(context.Background(), id, req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
		return
	} else {
		c.JSON(http.StatusOK, response.Success(nil))
	}
}

func GetUserOrderList(c *gin.Context) {
	userID := c.Param("user_id")
	// 转为int
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	mgr := order.GetOrderManager()
	if list, err := mgr.GetUserOrderList(context.Background(), id); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
		return
	} else {
		c.JSON(http.StatusOK, response.Success(list))
	}
}
