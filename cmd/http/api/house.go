package api

import (
	"context"
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/common/response"
	"github.com/cuixiaojun001/LinkHome/library/utils"
	"github.com/cuixiaojun001/LinkHome/services/house"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListHomeHouseInfo(c *gin.Context) {
	// 获取query参数city
	city := c.Query("city")
	mgr := house.GetHouseManager()
	if data, err := mgr.HomeHouseInfo(city); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(data))
	}
}

func PublishHouse(c *gin.Context) {
	var req house.PublishHouseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	mgr := house.GetHouseManager()
	if data, err := mgr.PublishHouse(req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(data))
	}
}

func GetHouseDetail(c *gin.Context) {
	id := c.Param("house_id")
	// 从request的headers中的Authorization获取值，去掉“Bearer ”前缀就是token
	token := c.GetHeader("Authorization")[7:]
	userID, err := utils.ParseJWTToken(token)
	if err != nil {
		logger.Errorw("parse jwt token failed", err)
		c.JSON(http.StatusOK, response.Unauthorized(err))
		return
	}
	houseID, _ := strconv.Atoi(id)
	mgr := house.GetHouseManager()
	if data, err := mgr.GetHouseDetail(context.Background(), houseID, userID); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(data))
	}
}

// ListHouse 获取房源列表信息
func ListHouse(c *gin.Context) {
	var req house.HouseListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	mgr := house.GetHouseManager()
	if data, err := mgr.HouseListInfo(req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(data))
	}
}

func GetAllHouseFacility(c *gin.Context) {
	mgr := house.GetHouseManager()
	if data, err := mgr.GetAllHouseFacility(); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(data))
	}
}

func GetRecommendHouseList(c *gin.Context) {
	token := c.GetHeader("Authorization")[7:]
	userID, err := utils.ParseJWTToken(token)
	if err != nil {
		logger.Errorw("parse jwt token failed", err)
		c.JSON(http.StatusOK, response.Unauthorized(err))
		return
	}

	req := &house.HouseListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	mgr := house.GetHouseManager()
	if data, err := mgr.GetRecommendHouseList(context.Background(), userID, req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(data))
	}
}

func UserHouseCollect(c *gin.Context) {
	req := &house.HouseCollectRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}

	//获取当前请求方法
	method := c.Request.Method
	mgr := house.GetHouseManager()
	if err := mgr.UserHouseCollect(context.Background(), method, req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(method))
	}
}

func GetUserHouseCollect(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}

	mgr := house.GetHouseManager()
	if data, err := mgr.GetUserHouseCollect(context.Background(), userID); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(data))
	}
}
