// cmd/main/factories/factories.go - Camada Principal - Factories
package factories

import (
	"fmt"
	"os"
	"strconv"

	"github.com/RafaelCava/kitkit-back-go/cmd/domain/models/user_models"
	"github.com/RafaelCava/kitkit-back-go/cmd/domain/usecases/user_usecase"
	"github.com/RafaelCava/kitkit-back-go/cmd/infra/repositories/user_repository"
	presentation "github.com/RafaelCava/kitkit-back-go/cmd/presentation/controllers/user_controller"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewDatabaseOpenConnection() error {
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT")) // don't forget to convert int since port is int type.
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo", host, user, pass, dbname, port)
	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db = con
	newMigrateModels()
	return err
}

func newMigrateModels() {
	// Migrar modelos para o banco de dados
	db.AutoMigrate(&user_models.User{})
}

func NewUserController() *presentation.UserController {
	userRepository := user_repository.NewGormUserRepository(db)
	userService := user_usecase.NewUserServiceImpl(userRepository)
	// Configurar controladores
	userController := presentation.NewUserController(userService)
	return userController
}
