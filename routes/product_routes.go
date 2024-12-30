package routes

import (
	"database/sql"
	"go-api/controller"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func InitRouteProduct(server *gin.Engine, connection *sql.DB) {

	repo := repository.NewProductRepository(connection)
	useCase := usecase.NewProductUseCase(repo)
	controller := controller.NewProductController(useCase)

	server.GET("/products", controller.GetProducts)
	server.GET("/product/:productId", controller.GetProductById)
	server.POST("/product", controller.CreateProduct)
}
