package main

import (
	"go-api/db"
	"go-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer dbConnection.Close()

	routes.InitRouteProduct(server, dbConnection)
	routes.InitRouteUser(server, dbConnection)

	server.Run(":8000")
}
