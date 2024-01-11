// main.go - Camada Principal
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/RafaelCava/kitkit-back-go/cmd/domain/user_models"
	"github.com/RafaelCava/kitkit-back-go/cmd/domain/user_usecase"
	"github.com/RafaelCava/kitkit-back-go/cmd/infra/user_repository"
	presentation "github.com/RafaelCava/kitkit-back-go/cmd/presentation/user_controller"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Configurar o Gorm
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT")) // don't forget to convert int since port is int type.
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo", host, user, pass, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	// Migrar modelos para o banco de dados
	db.AutoMigrate(&user_models.User{})

	// Configurar dependências
	userRepository := user_repository.NewGormUserRepository(db)
	userService := user_usecase.NewUserServiceImpl(userRepository)

	// Configurar roteador Gin
	router := gin.Default()

	// Configurar controladores
	userController := presentation.NewUserController(userService)
	userController.RegisterRoutes(router.Group("/user"))

	// Executar o servidor
	router.Run(":3000")
}
