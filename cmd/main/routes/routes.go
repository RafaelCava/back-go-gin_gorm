package routes

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// Run will start the server
func Run() {
	getRoutes()
	router.Run(":3000")
}

func getRoutes() {
	apiPrefix := router.Group("/api")
	addUserRoutes(apiPrefix)
}
