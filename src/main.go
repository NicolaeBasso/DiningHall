package main

import (
	"dininghall/src/configs"
	"dininghall/src/controllers"
	"dininghall/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()

	configs.SetupENV()
	services.InitCoreService()
	controllers.SetupController(router)
	// services.GenerateOrders(constants.GeneratedOrdersCount)

	router.Run(":4005")
}
