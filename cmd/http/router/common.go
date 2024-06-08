package router

import (
	"github.com/cuixiaojun001/LinkHome/cmd/http/api"
	"github.com/cuixiaojun001/LinkHome/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCommonAPI(engine *gin.Engine) {
	g := engine.Group("/api/v1/common").Use(middleware.Cors())
	{
		g.GET("/areas", api.AreaInfo)  // 获取省市区信息
		g.POST("/upload/", api.Upload) // 上传文件
		g.POST("/news", api.GetNews)   // 获取公告资讯列表
	}
}
