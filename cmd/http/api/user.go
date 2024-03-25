package api

import (
	"context"
	"github.com/cuixiaojun001/linkhome/common/logger"
	"github.com/cuixiaojun001/linkhome/common/response"
	"github.com/cuixiaojun001/linkhome/services/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {
	var req user.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	logger.Debugw("Login", "account", req.Account, "password", req.Password)
	if item, err := user.Login(context.TODO(), req); err != nil {
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
	if err := user.SendSmsCode(context.TODO(), mobile); err != nil {
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
	logger.Debugw("Register", "username", req.Username, "mobile", req.Mobile, "password", req.Password, "sms_code", req.SmsCode)
	if item, err := user.Register(context.TODO(), req); err != nil {
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
	if item, err := user.PwdChange(context.TODO(), userID, req); err != nil {
		c.JSON(http.StatusOK, response.BusinessException(err))
	} else {
		c.JSON(http.StatusOK, response.Success(map[string]string{"token": item.Token, "refresh_token": item.RefreshToken}))
	}
}

func UserProfile(c *gin.Context) {
	userID := c.Param("user_id")
	if item, err := user.Profile(context.TODO(), userID); err != nil {
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
