package router

import (
	"github.com/cuixiaojun001/LinkHome/cmd/http/api"
	"github.com/cuixiaojun001/LinkHome/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPaymentAPI(engine *gin.Engine) {
	g := engine.Group("/api/v1/payment").Use(middleware.Cors())
	{
		g.POST("/alipay/orders/:order_id", api.AliPayOrder) // 支付宝支付
		g.GET("/alipay/callback", api.AliPayCallback)       // 支付宝支付回调
	}
}
