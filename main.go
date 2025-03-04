package main

import (
	"github.com/gin-gonic/gin"
	employeeInfra "github.com/martinezgomez34/abarrote/src/employee/infraestructure"
	productsInfra "github.com/martinezgomez34/abarrote/src/products/infraestructure"
	
)

func main() {
	router := gin.Default()
	productsInfra.RegisterProductRoutes(router)
	employeeInfra.RegisterEmployeeRoutes(router)
	router.Run(":8080")
}