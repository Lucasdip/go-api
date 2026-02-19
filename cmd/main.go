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


// 1. Criamos o repositório e guardamos na variável 'ProductRepository'
ProductRepository := repository.NewProductRepository(dbConnection)

// 2. Passamos ESSA VARIÁVEL para o UseCase (sem repetir tipos ou nomes extras)
ProductUsecase := usecase.NewProductUseCase(ProductRepository)

// 3. Passamos o UseCase para o Controller
productController := controller.NewProductController(ProductUsecase)

server.GET("/ping", func (ctx *gin.Context)  {

	ctx.JSON(200, gin.H{
		"message": "pong",
	})
	
})

server.GET("/products", productController.GetProducts)

server.POST("/product", productController.CreateProduct)

server.GET("/product/:productId", productController.GetProductById)

server.Run(":8000")

}