package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func addDocsRoutes(rg *gin.RouterGroup) {
	rg.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
