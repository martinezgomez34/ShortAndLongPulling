package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/products/application"
	"github.com/martinezgomez34/abarrote/src/products/domain"
)

type EditProductController struct {
	service *application.EditProductUseCase
}

func NewEditProductController(service *application.EditProductUseCase) *EditProductController {
	return &EditProductController{service: service}
}

func (pc *EditProductController) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = int32(id)
	if err := pc.service.Update(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	application.NotifyProductAction()
	
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}