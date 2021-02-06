package main

import (
	"github.com/gin-gonic/gin"
	"github.com/morscino/teady-2-app/routes"
)

func main() {
	server := gin.Default()

	//All routes
	routes.HomeRoutes(server)
	routes.UserRoutes(server)

	server.Run(":9090")
}
