package router

import (
	"github.com/cuixiaojun001/linkhome/cmd/http/api"
	"github.com/gin-gonic/gin"
)

func RegisterUserAPI(engine *gin.Engine) {

	g := engine.Group("/api/v1/user")
	{
		// 用户登陆
		g.POST("/login", api.UserLogin)

		// 用户注册
		g.POST("/register", api.UserRegister)
		g.GET("/verify/username/:username", api.UsernameVerify)
		g.GET("/verify/mobile/:mobile", api.MobileVerify)
		g.GET("/sms_code/:mobile", api.SendSmsCode)

		// 用户信息
		g.PUT("/:user_id/pwd_change", api.UserPwdChange)
		g.GET("/profile/:user_id", api.UserProfile)
		g.PUT("/profile/:user_id", api.UserProfileUpdate)
	}
}
