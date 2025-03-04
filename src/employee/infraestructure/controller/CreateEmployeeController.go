package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/employee/application"
	"github.com/martinezgomez34/abarrote/src/employee/domain"
)

type CreateEmployeeController struct {
	employeeUseCase *application.CreateEmployeeUseCase
	notifyEmployeeCreated *application.NotifyEmployeeCreatedUseCase
}

func NewEmployee(uc *application.CreateEmployeeUseCase, notify *application.NotifyEmployeeCreatedUseCase) *CreateEmployeeController {
	return &CreateEmployeeController{employeeUseCase: uc, notifyEmployeeCreated: notify}
}

func (c *CreateEmployeeController) Save(ctx *gin.Context) {
	var employee domain.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := c.employeeUseCase.Create(&employee); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.notifyEmployeeCreated.NotifyEmployeeCreated(&employee)

	ctx.JSON(http.StatusCreated, gin.H{"message": "Employee created successfully"})
}
