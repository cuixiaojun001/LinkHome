package router

import (
	"github.com/cuixiaojun001/linkhome/cmd/http/api"
	"github.com/cuixiaojun001/linkhome/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCommonAPI(engine *gin.Engine) {
	g := engine.Group("/api/v1/common").Use(middleware.Cors())
	{
		// 用户登陆
		g.GET("/areas", api.AreaInfo)  // 获取省市区信息
		g.POST("/upload/", api.Upload) // 上传文件
	}
}
