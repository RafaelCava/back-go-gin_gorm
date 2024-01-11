package routes

import (
	"github.com/RafaelCava/kitkit-back-go/cmd/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.POST("/create", func(c *gin.Context) {
		controllers.AddUser(c)
	})
}
