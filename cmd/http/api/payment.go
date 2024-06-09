package api

import (
	"github.com/cuixiaojun001/LinkHome/common/response"
	"github.com/cuixiaojun001/LinkHome/services/payment"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AliPayOrder(c *gin.Context) {
	orderIDStr := c.Param("order_id")
	orderID, _ := strconv.Atoi(orderIDStr)

	req := &payment.OrderPaymentRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	if resp, err := payment.AliPayOrder(orderID, req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		// 重定向到http://localhost:6868/order.html
		c.JSON(http.StatusOK, response.Success(resp))
	}
}
