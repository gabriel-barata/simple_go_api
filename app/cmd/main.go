package main

import (
	"simple-go-api/app/controller"
	"simple-go-api/app/db"
	"simple-go-api/app/repository"
	"simple-go-api/app/usecase"
	"simple-go-api/app/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnvFile()

	server := gin.Default()
	dbConnection, _ := db.ConnectDB()

	// Repository layer
	productRepository := repository.NewProductRepository(dbConnection)
	// UseCase Layer
	productUsecase := usecase.NewProductUseCase(productRepository)
	// Controller Layer
	productController := controller.NewProductController(productUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.InsertProduct)
	server.GET("/product/:productId", productController.GetProductById)

	server.Run(":8000")
}
