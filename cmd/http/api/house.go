package api

import (
	"github.com/cuixiaojun001/linkhome/common/logger"
	"github.com/cuixiaojun001/linkhome/common/response"
	"github.com/cuixiaojun001/linkhome/services/house"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	if data, err := house.HouseListInfo(req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(data))
	}
}
