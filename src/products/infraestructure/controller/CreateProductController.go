package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/products/application"
	"github.com/martinezgomez34/abarrote/src/products/domain"
)

type CreateProductController struct {
	service *application.CreateProductUseCase
}

func NewProductController(service *application.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{service: service}
}

func (pc *CreateProductController) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := pc.service.Create(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	application.NotifyProductCreated()
	
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}