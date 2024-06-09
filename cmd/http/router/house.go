package router

import (
	"github.com/cuixiaojun001/LinkHome/cmd/http/api"
	"github.com/cuixiaojun001/LinkHome/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterHouseAPI(engine *gin.Engine) {
	g := engine.Group("/api/v1/house").Use(middleware.Cors())
	{
		g.GET("/home_houses", api.ListHomeHouseInfo)              // 获取首页房源信息
		g.POST("/houses", api.ListHouse)                          //获取房源列表信息
		g.GET("/houses/:house_id", api.GetHouseDetail)            //获取房源详情信息
		g.POST("/publish", api.PublishHouse)                      //发布房源信息
		g.GET("/facilities", api.GetAllHouseFacility)             //获取全部房屋设施信息
		g.POST("/recommend", api.GetRecommendHouseList)           //获取推荐房源信息
		g.POST("/user_collects", api.UserHouseCollect)            // 收藏房源信息
		g.DELETE("/user_collects", api.UserHouseCollect)          // 收藏房源信息
		g.GET("/user_collects/:user_id", api.GetUserHouseCollect) // 获取用户收藏房源信息
	}
}
