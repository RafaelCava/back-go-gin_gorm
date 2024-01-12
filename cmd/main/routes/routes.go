package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	apiPrefix.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
