package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type productController struct {

	ProductUsecase usecase.ProductUsecase

}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{ProductUsecase: usecase}
}

func (p *productController) GetProducts(ctx *gin.Context) {


	products, err := p.ProductUsecase.GetProducts()
	if (err != nil) {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}

	
func (p *productController) CreateProduct(ctx *gin.Context){

	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.ProductUsecase.CreateProduct(product)
		
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
    // 1. Pega o ID da URL (ele vem como string)
    id := ctx.Param("productId")
    if id == "" {
        ctx.JSON(400, gin.H{"error": "ID do produto é obrigatório"})
        return
    }

    // 2. Converte a string para int (opcional, mas recomendado para o banco)
    idInt, err := strconv.Atoi(id)
    if err != nil {
        ctx.JSON(400, gin.H{"error": "ID precisa ser um número"})
        return
    }

    // 3. Chama o Usecase
    product, err := p.ProductUsecase.GetProductById(idInt)
    if err != nil {
        ctx.JSON(404, gin.H{"error": "Produto não encontrado"})
        return
    }

    ctx.JSON(200, product)
}

