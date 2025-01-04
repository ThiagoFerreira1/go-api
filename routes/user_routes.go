package routes

import (
	"database/sql"
	"go-api/controller"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func InitRouteUser(server *gin.Engine, connection *sql.DB) {
	repo := repository.NewUserRepository(connection)
	useCase := usecase.NewUserUseCase(repo)
	controller := controller.NewUserController(useCase)

	server.POST("/user", controller.CreateUser)
}
