package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(engine *gin.Engine) {
	// RegisterFrontend(engine)
	RegisterHouseAPI(engine)
	RegisterUserAPI(engine)
	RegisterCommonAPI(engine)
	RegisterOrderAPI(engine)
}

// GinHandler 将http.HandlerFunc转为gin.HandlerFunc
func GinHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
