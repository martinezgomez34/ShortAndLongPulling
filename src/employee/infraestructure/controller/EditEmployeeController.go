package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/employee/application"
	"github.com/martinezgomez34/abarrote/src/employee/domain"
)

type EditEmployeeController struct {
	employeeUseCase       *application.EditEmployeeUseCase
	notifyEmployeeAction  *application.NotifyEmployeeActionUseCase
}

func EditEmployee(uc *application.EditEmployeeUseCase, notify *application.NotifyEmployeeActionUseCase) *EditEmployeeController {
	return &EditEmployeeController{employeeUseCase: uc, notifyEmployeeAction: notify}
}

func (pc *EditEmployeeController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var employee domain.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	employee.ID = int32(id)
	if err := pc.employeeUseCase.Update(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pc.notifyEmployeeAction.RegisterAction(employee.ID, "updated")

	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
}
