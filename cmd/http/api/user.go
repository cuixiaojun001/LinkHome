package api

import (
	"context"
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"net/http"
	"strconv"

	"github.com/cuixiaojun001/LinkHome/common/response"
	"github.com/cuixiaojun001/LinkHome/services/user"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var req user.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	mgr := user.GetUsereManager()
	if item, err := mgr.Login(context.TODO(), req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
		return
	} else {
		c.JSON(http.StatusOK, response.Success(map[string]string{"token": item.Token, "refresh_token": item.RefreshToken}))
	}
}

func UsernameVerify(c *gin.Context) {
	username := c.Param("username")
	if user.IsUsernameExist(username) {
		c.JSON(http.StatusOK, response.Failed(map[string]bool{"verify_result": true}))
	} else {
		c.JSON(http.StatusOK, response.Success(map[string]bool{"verify_result": false}))
	}
}

func MobileVerify(c *gin.Context) {
	mobile := c.Param("mobile")
	if user.IsMobileExist(mobile) {
		c.JSON(http.StatusOK, response.Failed(map[string]bool{"verify_result": true}))
	} else {
		c.JSON(http.StatusOK, response.Success(map[string]bool{"verify_result": false}))
	}
}

func SendSmsCode(c *gin.Context) {
	mobile := c.Param("mobile")
	mgr := user.GetUsereManager()
	if err := mgr.SendSmsCode(context.TODO(), mobile); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(nil))
	}
}

func UserRegister(c *gin.Context) {
	var req user.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}

	mgr := user.GetUsereManager()
	if item, err := mgr.Register(context.TODO(), req); err != nil {
		c.JSON(http.StatusOK, response.BusinessException(err))
	} else {
		c.JSON(http.StatusOK, response.Success(map[string]string{"token": item.Token, "refresh_token": item.RefreshToken}))
	}
}

func UserPwdChange(c *gin.Context) {
	userID := c.Param("user_id")
	var req user.PwdChangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}

	mgr := user.GetUsereManager()
	if item, err := mgr.PwdChange(context.TODO(), userID, req); err != nil {
		c.JSON(http.StatusOK, response.BusinessException(err))
	} else {
		c.JSON(http.StatusOK, response.Success(map[string]string{"token": item.Token, "refresh_token": item.RefreshToken}))
	}
}

func UserProfile(c *gin.Context) {
	userID := c.Param("user_id")
	// 转为int
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	mgr := user.GetUsereManager()
	if item, err := mgr.Profile(context.TODO(), id); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(item))
	}
}

func UserProfileUpdate(c *gin.Context) {
	userID := c.Param("user_id")
	var req user.ProfileUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	if item, err := user.ProfileUpdate(context.TODO(), userID, req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(item))
	}
}

func PublishOrUpdateUserRentalDemand(c *gin.Context) {
	userIDStr := c.Param("user_id")
	// 将字符串转换为int
	id, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	var req user.RentalDemandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}

	mgr := user.GetUsereManager()
	if err := mgr.PublishOrUpdateRentalDemand(context.TODO(), id, req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(nil))
	}
}

func UserRealNameAuth(c *gin.Context) {
	var req user.UserRealNameAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}

	mgr := user.GetUsereManager()
	if resp, err := mgr.UserRealNameAuth(context.TODO(), req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(resp))
	}
}

func GetUserRentalDemands(c *gin.Context) {
	req := &user.RentalDemandListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}

	mgr := user.GetUsereManager()
	if resp, err := mgr.GetUserRentalDemands(context.TODO(), req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(resp))
	}
}

func GetRentalDemandDetail(c *gin.Context) {
	idStr := c.Param("demand_id")
	logger.Debugw("GetRentalDemandDetail", "idStr", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}

	mgr := user.GetUsereManager()
	logger.Debugw("GetRentalDemandDetail", "id", id)
	if resp, err := mgr.GetRentalDemandDetail(context.TODO(), id); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(resp))
	}
}
