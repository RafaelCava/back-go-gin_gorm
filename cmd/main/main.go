// cmd/main/main.go - Camada Principal
package main

import (
	"github.com/RafaelCava/kitkit-back-go/cmd/main/factories"
	"github.com/RafaelCava/kitkit-back-go/cmd/main/routes"
)

func main() {
	err := factories.NewDatabaseOpenConnection()
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	// Executar o servidor
	routes.Run()
}
