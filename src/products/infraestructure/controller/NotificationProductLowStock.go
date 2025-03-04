package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/products/application"
)

type NotifyProductActionController struct {
	service *application.NotifyProductActionUseCase
}

func NewNotifyProductActionController(service *application.NotifyProductActionUseCase) *NotifyProductActionController {
	return &NotifyProductActionController{service: service}
}

func (pc *NotifyProductActionController) WaitForAction(c *gin.Context) {
	products, err := pc.service.WaitForAction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if products == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No hay cambios"})
		return
	}

	var messages []string
	for _, product := range products {
		messages = append(messages, "Se ha actualizado el producto: "+product.Name+" ID: "+strconv.Itoa(int(product.ID)))
	}

	c.JSON(http.StatusOK, gin.H{"message": messages})
}
