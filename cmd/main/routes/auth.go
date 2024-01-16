package routes

import (
	"github.com/RafaelCava/kitkit-back-go/cmd/main/factories"
	"github.com/gin-gonic/gin"
)

func addAuthRoutes(rg *gin.RouterGroup) {
	authPrefix := rg.Group("/auth")
	factories.NewAuthController().RegisterRoutes(authPrefix)
}
