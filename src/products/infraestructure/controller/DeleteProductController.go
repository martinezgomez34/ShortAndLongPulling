package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/products/application"
)

type DeleteProductController struct {
	service *application.DeleteProductUseCase
}

func NewDeleteProductController(service *application.DeleteProductUseCase) *DeleteProductController {
	return &DeleteProductController{service: service}
}

func (pc *DeleteProductController) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := pc.service.Delete(int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	application.NotifyProductAction()
	
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}