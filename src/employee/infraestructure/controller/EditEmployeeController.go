package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/employee/application"
	"github.com/martinezgomez34/abarrote/src/employee/domain"
)

type EditEmployeeController struct {
	employeeUseCase *application.EditEmployeeUseCase
}

func EditEmployee(uc *application.EditEmployeeUseCase) *EditEmployeeController {
	return &EditEmployeeController{employeeUseCase: uc}
}

func (pc *EditEmployeeController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var enployee domain.Employee
	if err := c.ShouldBindJSON(&enployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	enployee.ID = int16(id)
	if err := pc.employeeUseCase.Update(&enployee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}