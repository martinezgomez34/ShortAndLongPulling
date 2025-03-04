package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/employee/application"
)

type DeleteEmployeeController struct {
	employeeUseCase *application.DeleteEmployeeUseCase
}

func DeleteEmployee(uc *application.DeleteEmployeeUseCase) *DeleteEmployeeController {
	return &DeleteEmployeeController{employeeUseCase: uc}
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
	ctx.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}