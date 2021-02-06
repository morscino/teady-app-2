package facade

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morscino/teady-2-app/handlers"
	"github.com/morscino/teady-2-app/service/userservice"
)

type UserFascade struct {
	ctx         context.Context
	userhandler handlers.UserHandler
}

func NewUserFacade(ctx context.Context, u handlers.UserHandler) *UserFascade {

	return &UserFascade{
		ctx:         ctx,
		userhandler: u,
	}
}

func (f UserFascade) CreateUser(c *gin.Context) {
	var input userservice.UserRegistrationData
	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := f.userhandler.CreateUser(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
