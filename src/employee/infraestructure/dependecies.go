package infraestructure

import (
	"github.com/gin-gonic/gin"
	MySQL "github.com/martinezgomez34/abarrote/src/core"
	"github.com/martinezgomez34/abarrote/src/employee/application"
	"github.com/martinezgomez34/abarrote/src/employee/infraestructure/controller"
	"github.com/martinezgomez34/abarrote/src/employee/infraestructure/rute"
)

func RegisterEmployeeRoutes(router *gin.Engine) {
	db := MySQL.GetDBPool()

	repo := NewSQL(db)

	createEmployeeUC := application.CreateEmployee(repo)
	listEmployeeUC := application.ListEmployee(repo)
	editEmployeeUC := application.EditEmployee(repo)
	deleteEmployeeUC := application.DeleteEmployee(repo)

	notifyEmployeeCreatedUC := application.NewNotifyEmployeeCreatedUC()  
	notifyEmployeeActionUC := application.NewNotifyEmployeeActionUC() 

	createEmployeeController := controller.NewEmployee(createEmployeeUC, notifyEmployeeCreatedUC)
	listEmployeeController := controller.ListEmployee(listEmployeeUC)
	editEmployeeController := controller.EditEmployee(editEmployeeUC, notifyEmployeeActionUC)
	deleteEmployeeController := controller.DeleteEmployee(deleteEmployeeUC, notifyEmployeeActionUC)


	notifyEmployeeCreatedController := controller.NewNotifyEmployeeCreatedController(notifyEmployeeCreatedUC)
	notifyEmployeeActionController := controller.NewNotifyEmployeeActionController(notifyEmployeeActionUC)

	rute.EmployeeRoutes(router, createEmployeeController, listEmployeeController, editEmployeeController, deleteEmployeeController, 
		notifyEmployeeCreatedController, notifyEmployeeActionController)
}
