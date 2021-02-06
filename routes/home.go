package routes

import (
	"github.com/gin-gonic/gin"
)

func HomeRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		//add content
		ctx.JSON(200, gin.H{
			"message": "we are live",
		})
	})
}
