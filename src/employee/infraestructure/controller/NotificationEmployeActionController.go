package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/employee/application"
)

type NotifyEmployeeActionController struct {
	service *application.NotifyEmployeeActionUseCase
}

func NewNotifyEmployeeActionController(service *application.NotifyEmployeeActionUseCase) *NotifyEmployeeActionController {
	return &NotifyEmployeeActionController{service: service}
}

func (pc *NotifyEmployeeActionController) GetEmployeeActions(c *gin.Context) {
	actions := pc.service.GetActions()
	if actions == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No hay acciones registradas"})
		return
	}

	var messages []string
	for _, action := range actions {
		messages = append(messages, "Acción: "+action.Action+" en el empleado con ID: "+strconv.Itoa(int(action.ID)))
	}

	c.JSON(http.StatusOK, gin.H{"actions": messages})
}
