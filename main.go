package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"github.com/morscino/teady-app-2/database/postgres"
	"github.com/morscino/teady-app-2/facade"
	"github.com/morscino/teady-app-2/handlers"
	"github.com/morscino/teady-app-2/models/usermodel"
	"github.com/morscino/teady-app-2/routes"
	"github.com/morscino/teady-app-2/service/userservice"
	"github.com/morscino/teady-app-2/utility/config"
	"github.com/morscino/teady-app-2/utility/log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Error("Failed to load env data: %v", err)
	}

	var config config.Config

	if err := envconfig.Process("", &config); err != nil {
		log.Error("Could not load configuration data: %v", err)
	}

	db := postgres.DbConnect(config.DB)
	//migrations
	var UserModel usermodel.User
	db.AutoMigrate(&UserModel)

	server := gin.Default()
	var ctx context.Context

	//All routes
	routes.HomeRoutes(server)

	UserRepository := userservice.NewUserService(db)
	UserHandler := handlers.NewUserHandler(UserRepository)
	UserFacade := *facade.NewUserFacade(ctx, UserHandler)
	u := routes.NewUserRoute(UserFacade)

	u.UserRoutes(server)

	server.Run(":" + config.App.Port)
}
