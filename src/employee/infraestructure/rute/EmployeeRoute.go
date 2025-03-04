package rute

import (
	"github.com/gin-gonic/gin"
	"github.com/martinezgomez34/abarrote/src/employee/infraestructure/controller"
)

func EmployeeRoutes(router *gin.Engine, createEmployeeController *controller.CreateEmployeeController,
	listEmployeeController *controller.ListEmployeeController,  editemployeeController *controller.EditEmployeeController,
	deleteEmployeeController *controller.DeleteEmployeeController, getByEmployeeController *controller.GetByIDEmployeeController) {
	employeeGroup := router.Group("/employees")
	{
		employeeGroup.POST("/", createEmployeeController.Save)
		employeeGroup.GET("/", listEmployeeController.GetAll)
		employeeGroup.GET("/:id", getByEmployeeController.GetEmployeeByID)
		employeeGroup.PUT("/:id", editemployeeController.Update)
		employeeGroup.DELETE("/:id", deleteEmployeeController.Delete)
	}
}