package router

import (
	"github.com/cuixiaojun001/LinkHome/cmd/http/api"
	"github.com/gin-gonic/gin"
)

func RegisterHouseAPI(engine *gin.Engine) {
	g := engine.Group("/api/v1/house")
	{
		g.GET("/home_houses", api.ListHomeHouseInfo)   // 获取首页房源信息
		g.POST("/houses", api.ListHouse)               //获取房源列表信息
		g.GET("/houses/:house_id", api.GetHouseDetail) //获取房源详情信息
		g.POST("/publish", api.PublishHouse)           //发布房源信息
		g.GET("/facilities", api.GetAllHouseFacility)  //获取全部房屋设施信息
	}
}
