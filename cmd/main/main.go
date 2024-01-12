// cmd/main/main.go - Camada Principal
package main

import (
	"github.com/RafaelCava/kitkit-back-go/cmd/domain/usecases/user_usecase"
	"github.com/RafaelCava/kitkit-back-go/cmd/infra/repositories/user_repository"
	"github.com/RafaelCava/kitkit-back-go/cmd/main/factories"
	presentation "github.com/RafaelCava/kitkit-back-go/cmd/presentation/controllers/user_controller"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := factories.NewDatabaseOpenConnection()
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	// Migrar modelos para o banco de dados
	factories.NewMigrateModels(db)

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
