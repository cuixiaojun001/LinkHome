package api

import (
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/common/response"
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
	houseID := c.Param("house_id")
	id, _ := strconv.Atoi(houseID)
	logger.Debugw("GetHouse", "houseId", id)
	mgr := house.GetHouseManager()
	if data, err := mgr.GetHouseDetail(id); err != nil {
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
