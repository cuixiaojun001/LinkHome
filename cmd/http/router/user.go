package router

import (
	"github.com/cuixiaojun001/linkhome/cmd/http/api"
	"github.com/gin-gonic/gin"
)

// RegisterUserAPI 用户模块
func RegisterUserAPI(engine *gin.Engine) {

	g := engine.Group("/api/v1/user")
	{
		// 用户登陆
		g.POST("/login", api.UserLogin)       // 用户登录
		g.POST("/register", api.UserRegister) // 用户注册

		g.GET("/verify/username/:username", api.UsernameVerify) // 用户名校验
		g.GET("/verify/mobile/:mobile", api.MobileVerify)       // 用户手机号校验
		g.GET("/sms_code/:mobile", api.SendSmsCode)             // 发送短信验证码

		// 用户信息
		g.PUT("/:user_id/pwd_change", api.UserPwdChange)  // 用户密码修改
		g.GET("/profile/:user_id", api.UserProfile)       // 获取用户详情
		g.PUT("/profile/:user_id", api.UserProfileUpdate) // 更新用户详情

		g.POST("/rental_demands/:user_id", api.PublishOrUpdateUserRentalDemand) // 发布或更新用户租房需求
		g.PUT("/rental_demands/:user_id", api.PublishOrUpdateUserRentalDemand)  // 发布或更新用户租房需求
	}
}
