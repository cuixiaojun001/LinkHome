package api

import (
	"github.com/cuixiaojun001/linkhome/common/logger"
	"github.com/cuixiaojun001/linkhome/common/response"
	"github.com/cuixiaojun001/linkhome/modules/house/model"
	"github.com/cuixiaojun001/linkhome/services/house"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ListHouseInfo 获取房屋信息列表
func ListHouseInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"errno":  0,
		"errmsg": "",
		"data": gin.H{
			"pageSize": 20,
			"total":    120,
			"list": []gin.H{
				{
					"surfaceImage":      "https://img.ljcdn.com/beike//beike/kemis/hhqezqklzl263.png",
					"houseStatus":       "1",
					"communityName":     "如视幸福新区123",
					"bizCircleName":     "测试西二旗",
					"houseDelCode":      "106104250671",
					"unitType":          "3-1-1-2",
					"areaSize":          180,
					"totalPrice":        2660000,
					"totalPriceStr":     "266万",
					"priceTrend":        0,
					"unitPrice":         14777,
					"floor":             0,
					"floorType":         "高",
					"totalFloor":        "5",
					"orientation":       "东南",
					"visitCount":        0,
					"followUp":          false,
					"createTime":        1683355706814,
					"maintainerId":      nil,
					"maintainerName":    nil,
					"inPoolTime":        "nkbpcywcph300",
					"inPoolReason":      "超期未维护",
					"overdueWarning":    nil,
					"inPoolRule":        nil,
					"inventoryScoreStr": nil,
					"inventoryTime":     nil,
					"inventoryTimeStr":  nil,
					"unInventoryLabel":  false,
					"sysScore":          nil,
					"voteNumStr":        nil,
					"tags": []string{
						"钥匙",
						"外网已呈现",
					},
					"tagSort": nil,
					"tagsSimplifyAndDetail": []gin.H{
						{
							"simple":   "钥",
							"complete": "钥匙",
							"type":     "usual_key",
							"key":      5,
						},
						{
							"simple":   "呈现",
							"complete": "外网已呈现",
							"type":     "appearance",
							"key":      22,
						},
					},
					"roleList":            []string{},
					"qualityScore":        nil,
					"exFieldData":         nil,
					"extendParam":         nil,
					"flag":                nil,
					"distance":            "未知",
					"subwayLineName":      "未知",
					"subwayName":          "未知",
					"geoDistance":         nil,
					"communityId":         "61rytamyvfqdy96",
					"ownerPhone1":         "",
					"ownerPhone2":         "",
					"homePhone":           "",
					"otherPhone":          "",
					"paymentMode":         "307500000001",
					"inventoryIsManager":  nil,
					"systemProb":          nil,
					"sysStars":            nil,
					"approBrokerUcid":     nil,
					"hasApproBroker":      false,
					"houseLastVisitTime":  nil,
					"realProspectingTime": nil,
					"companyCode":         nil,
					"brand":               0,
					"holderUcid":          nil,
					"showRemind":          false,
					"remindString":        nil,
					"vrStatus":            0,
					"claimWarning":        true,
					"fbExpoId":            nil,
					"signTotalPriceStr":   nil,
					"signPriceUnit":       nil,
					"signBrokerId":        nil,
					"signBrokerName":      nil,
					"signTime":            nil,
					"signPeriod":          0,
					"delType":             1,
					"quickTags":           []string{},
					"hurrySale":           false,
				},
			},
		},
	})
}

func ListHomeHouseInfo(c *gin.Context) {
	// 获取query参数city
	city := c.Query("city")
	if data, err := house.HomeHouseInfo(city); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(data))
	}
}

// ListHouse 获取房源列表信息
func ListHouse(c *gin.Context) {
	// 获取post参数
	var req model.HouseListRequest
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

func PublishHouse(c *gin.Context) {
	// 获取post参数
	var req model.PublishHouseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	if data, err := house.PublishHouse(req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(data))
	}
}

func GetHouse(c *gin.Context) {
	houseID := c.Param("house_id")
	id, _ := strconv.Atoi(houseID)
	logger.Debugw("GetHouse", "houseId", id)
	if data, err := house.GetHouse(id); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(data))
	}
}

func GetAllHouseFacility(c *gin.Context) {
	if data, err := house.GetAllHouseFacility(); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(data))
	}
}
