package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Camada de usecases
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)

	// Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	// Mapear a rota
	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
