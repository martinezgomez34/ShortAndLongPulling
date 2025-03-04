package controllers

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/products/application"
)

type NotifyProductCreatedController struct {
	service *application.NotifyProductCreatedUseCase
}

func NewNotifyProductCreatedController(service *application.NotifyProductCreatedUseCase) *NotifyProductCreatedController {
	return &NotifyProductCreatedController{service: service}
}

func (pc *NotifyProductCreatedController) WaitForNewProduct(c *gin.Context) {
	products, err := pc.service.WaitForNewProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if products == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No hay cambios"})
		return
	}

	// Formatear la respuesta
	var messages []string
	for _, product := range products {
		messages = append(messages, "Se ha agregado un nuevo producto: "+product.Name)
	}

	c.JSON(http.StatusOK, gin.H{"message": messages})
}
