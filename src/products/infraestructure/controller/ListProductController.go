package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/products/application"
)

type ListProductController struct {
	service *application.ListProductUseCase
}

func NewListProductController(service *application.ListProductUseCase) *ListProductController {
	return &ListProductController{service: service}
}

func (pc *ListProductController) GetAllProducts(c *gin.Context) {
	products, err := pc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}