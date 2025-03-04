package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/products/application"
)

type NotifyLowStockController struct {
	service *application.NotifyLowStockUseCase
}

func NewNotifyLowStockController(service *application.NotifyLowStockUseCase) *NotifyLowStockController {
	return &NotifyLowStockController{service: service}
}

func (pc *NotifyLowStockController) GetLowStockProducts(c *gin.Context) {
	products, err := pc.service.GetLowStockProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var messages []string
	for _, product := range products {
		messages = append(messages, "El producto "+product.Name+" está por agotarse, quedan "+strconv.Itoa(int(product.Amount))+" en stock ID: "+strconv.Itoa(int(product.ID)))
	}

	c.JSON(http.StatusOK, gin.H{"message": messages})
}