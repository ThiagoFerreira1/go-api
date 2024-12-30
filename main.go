package main

import (
	"go-api/db"
	"go-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o servidor
	server := gin.Default()

	// Conecta ao banco de dados
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer dbConnection.Close()

	// Inicializa as rotas
	routes.InitRouteProduct(server, dbConnection)

	// Inicia o servidor
	server.Run(":8000")
}
