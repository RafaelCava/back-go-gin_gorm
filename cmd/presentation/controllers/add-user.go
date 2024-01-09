package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RafaelCava/kitkit-back-go/cmd/domain/models"
	"github.com/RafaelCava/kitkit-back-go/cmd/infra/database"
	"github.com/gin-gonic/gin"
)

func AddUser(ctx *gin.Context) {
	body := models.User{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, "User is not defined")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Bad Input")
		return
	}

	result := database.Db.Create(&models.User{Username: body.Username, Password: body.Password})
	if result.Error != nil {
		fmt.Println(result.Error)
		ctx.AbortWithStatusJSON(400, "Couldn't create the new user.")
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "User created successfully!",
			"rows":    result.RowsAffected,
		})
	}

}
