package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/products/application"
)

type GetByIDProductController struct {
	service *application.GetByIDProductUseCase
}

func NewGetByIDProductController(service *application.GetByIDProductUseCase) *GetByIDProductController {
	return &GetByIDProductController{service: service}
}

func (pc *GetByIDProductController) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := pc.service.GetByID(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}