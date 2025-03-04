package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/employee/application"
)

type ListEmployeeController struct {
	employeeUseCase *application.ListEmployeeUseCase
}

func ListEmployee(uc *application.ListEmployeeUseCase) *ListEmployeeController {
	return &ListEmployeeController{employeeUseCase: uc}
}

func (c *ListEmployeeController) GetAll(ctx *gin.Context) {
	employees, err := c.employeeUseCase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, employees)
}
