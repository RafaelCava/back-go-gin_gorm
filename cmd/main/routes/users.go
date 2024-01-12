package routes

import (
	"github.com/RafaelCava/kitkit-back-go/cmd/main/factories"
	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	factories.NewUserController().RegisterRoutes(users)
}
