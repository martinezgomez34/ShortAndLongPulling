package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/employee/application"
)

type DeleteEmployeeController struct {
	employeeUseCase       *application.DeleteEmployeeUseCase
	notifyEmployeeAction  *application.NotifyEmployeeActionUseCase
}

func DeleteEmployee(uc *application.DeleteEmployeeUseCase, notify *application.NotifyEmployeeActionUseCase) *DeleteEmployeeController {
	return &DeleteEmployeeController{
		employeeUseCase:      uc,
		notifyEmployeeAction: notify,
	}
}

func (c *DeleteEmployeeController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.employeeUseCase.Delete(int32(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.notifyEmployeeAction.RegisterAction(int32(id), "deleted")

	ctx.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
