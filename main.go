package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/morscino/teady-2-app/facade"
	"github.com/morscino/teady-2-app/handlers"
	"github.com/morscino/teady-2-app/routes"
	"github.com/morscino/teady-2-app/service/userservice"
)

func main() {
	server := gin.Default()
	var ctx context.Context

	//All routes
	routes.HomeRoutes(server)
	UserRepository := userservice.NewUserService()
	UserHandler := handlers.NewUserHandler(UserRepository)
	UserFacade := *facade.NewUserFacade(ctx, UserHandler)
	u := routes.NewUserRoute(UserFacade)

	u.UserRoutes(server)

	server.Run(":9090")
}
