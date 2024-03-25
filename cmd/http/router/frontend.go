package router

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func RegisterFrontend(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		tmpl, err := template.ParseFiles("front/index.html")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = tmpl.Execute(c.Writer, nil)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
	})
}
