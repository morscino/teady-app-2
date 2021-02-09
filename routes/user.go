package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/morscino/teady-app-2/facade"
)

type UserRoute struct {
	UserFacacde facade.UserFascade
}

func NewUserRoute(f facade.UserFascade) *UserRoute {
	return &UserRoute{UserFacacde: f}
}
func (u UserRoute) UserRoutes(router *gin.Engine) {
	UserGroup := router.Group("/user")
	{
		UserGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "users connected",
			})
		})

		UserGroup.POST("/create-user", u.UserFacacde.CreateUser)
	}
}
