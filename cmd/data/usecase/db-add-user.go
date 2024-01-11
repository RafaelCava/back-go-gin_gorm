package data

import (
	"github.com/RafaelCava/kitkit-back-go/cmd/domain/models"
	"github.com/RafaelCava/kitkit-back-go/cmd/infra/database"
)

func AddUser(username, password string) (int, error) {
	result := database.Db.Create(&models.User{Username: username, Password: password})
	return int(result.RowsAffected), result.Error
}
