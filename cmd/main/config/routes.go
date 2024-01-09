package config

import (
	"github.com/RafaelCava/kitkit-back-go/cmd/main/routes"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// Run will start the server
func Init() {
	getRoutes()
	router.Run(":3000")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	prefixGroup := router.Group("/api")
	routes.AddUserRoutes(prefixGroup)
}
