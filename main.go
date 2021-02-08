package main

import (
	"context"

	"github.com/morscino/teady-2-app/utility/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/morscino/teady-2-app/database/postgres"
	"github.com/morscino/teady-2-app/facade"
	"github.com/morscino/teady-2-app/handlers"
	"github.com/morscino/teady-2-app/models/usermodel"
	"github.com/morscino/teady-2-app/routes"
	"github.com/morscino/teady-2-app/service/userservice"
)

func main() {
	godotenv.Load()

	var c config.Config

	envconfig.Process("", &c)

	server := gin.Default()
	var ctx context.Context
	var UserModel usermodel.User

	db := postgres.DbConnect(c.DB)

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
