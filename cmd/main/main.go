// cmd/main/main.go - Camada Principal
package main

import (
	"github.com/RafaelCava/kitkit-back-go/cmd/main/factories"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := factories.NewDatabaseOpenConnection()
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	// Migrar modelos para o banco de dados
	factories.NewMigrateModels(db)

	// Configurar roteador Gin
	router := gin.Default()

	// Configurar controladores
	userController := factories.NewUserController(db)
	userController.RegisterRoutes(router.Group("/user"))

	// Executar o servidor
	router.Run(":3000")
}
