// cmd/main/main.go - Camada Principal
package main

import (
	_ "github.com/RafaelCava/kitkit-back-go/cmd/main/docs"

	"github.com/RafaelCava/kitkit-back-go/cmd/main/factories"
	"github.com/RafaelCava/kitkit-back-go/cmd/main/routes"
)

//	@title			KitKit back golang
//	@version		1.0
//	@description	Golang server
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Rafael Cavalcante
//	@contact.email	rafael.cavalcante@tallos.com.br

//	@host		localhost:3000
//	@BasePath	/api

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	err := factories.NewDatabaseOpenConnection()
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	// Executar o servidor
	routes.Run()
}
