package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/morscino/teady-2-app/database/postgres"
	"github.com/morscino/teady-2-app/facade"
	"github.com/morscino/teady-2-app/handlers"
	"github.com/morscino/teady-2-app/models/usermodel"
	"github.com/morscino/teady-2-app/routes"
	"github.com/morscino/teady-2-app/service/userservice"
)

func main() {
	server := gin.Default()
	var ctx context.Context
	var UserModel usermodel.User

	db := postgres.DbConnect()

	//migrations
	db.AutoMigrate(&UserModel)

	//All routes
	routes.HomeRoutes(server)
	UserRepository := userservice.NewUserService(db)
	UserHandler := handlers.NewUserHandler(UserRepository)
	UserFacade := *facade.NewUserFacade(ctx, UserHandler)
	u := routes.NewUserRoute(UserFacade)

	u.UserRoutes(server)

	server.Run(":9090")
}
