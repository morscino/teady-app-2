package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	UserGroup := router.Group("/user")
	{
		UserGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "users connected",
			})
		})
	}
}
