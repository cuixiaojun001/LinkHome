package router

import (
	"github.com/cuixiaojun001/LinkHome/cmd/http/api"
	"github.com/gin-gonic/gin"
)

// RegisterAdminAPI 后台管理模块
func RegisterAdminAPI(engine *gin.Engine) {
	g := engine.Group("/api/v1/admin")
	{
		g.POST("/user/users", api.UserList)
		g.PUT("/houses/:house_id", api.Upload)
		g.POST("/orders", api.OrderList)
	}
}
