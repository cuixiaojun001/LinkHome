package router

import (
	"github.com/cuixiaojun001/linkhome/cmd/http/api"
	"github.com/gin-gonic/gin"
)

func RegisterOrderAPI(engine *gin.Engine) {
	g := engine.Group("/api/v1/order") // .Use(middleware.Cors())
	{
		g.POST("/orders/:user_id", api.CreateOrder)     // 创建租房订单
		g.GET("/orders/:user_id", api.GetUserOrderList) // 获取用户租房订单
	}
}
