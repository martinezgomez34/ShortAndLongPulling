package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/employee/application"
)

type NotifyEmployeeCreatedController struct {
	service *application.NotifyEmployeeCreatedUseCase
}

func NewNotifyEmployeeCreatedController(service *application.NotifyEmployeeCreatedUseCase) *NotifyEmployeeCreatedController {
	return &NotifyEmployeeCreatedController{service: service}
}

func (pc *NotifyEmployeeCreatedController) WaitForNewEmployee(c *gin.Context) {
	employee, err := pc.service.WaitForNewEmployee()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if employee == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No hay empleados nuevos"})
		return
	}

	// Notificamos que se ha creado un nuevo empleado
	c.JSON(http.StatusOK, gin.H{"message": "Nuevo empleado creado: " + employee.FirstName})
}
