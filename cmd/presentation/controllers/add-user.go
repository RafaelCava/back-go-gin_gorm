package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	data "github.com/RafaelCava/kitkit-back-go/cmd/data/usecase"
	"github.com/RafaelCava/kitkit-back-go/cmd/domain/models"
	"github.com/gin-gonic/gin"
)

func AddUser(ctx *gin.Context, addUser ) {
	body := models.User{}
	dataRaw, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, "User is not defined")
		return
	}
	err = json.Unmarshal(dataRaw, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Bad Input")
		return
	}
	rowsAffected, dataError := data.AddUser(body.Username, body.Password)
	if dataError != nil {
		fmt.Println(dataError)
		ctx.AbortWithStatusJSON(400, "Couldn't create the new user.")
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "User created successfully!",
			"rows":    rowsAffected,
		})
	}

}
