package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.GET("/all", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "all users",
		})
	})
}
