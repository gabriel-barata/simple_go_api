package controller

import (
	"net/http"
	"simple-go-api/app/models"

	"github.com/gin-gonic/gin"
)

type productController struct {
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products := []models.Product{
		{
			ID:    1,
			Name:  "Molho de tomate",
			Price: 13.56,
		},
	}

	ctx.JSON(http.StatusOK, products)

}
