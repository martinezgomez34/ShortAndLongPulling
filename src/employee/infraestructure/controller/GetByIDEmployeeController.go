package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/employee/application"
)

type GetByIDEmployeeController struct {
	service *application.GetByIDEmployeeUseCase
}

func NewGetByIDProductController(service *application.GetByIDEmployeeUseCase) *GetByIDEmployeeController {
	return &GetByIDEmployeeController{service: service}
}

func (pc *GetByIDEmployeeController) GetEmployeeByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := pc.service.GetByID(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}