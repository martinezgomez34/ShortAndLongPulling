package main

import (
	"github.com/gin-gonic/gin"
	productsInfra "github.com/martinezgomez34/abarrote/src/products/infraestructure"
)

func main() {
	router := gin.Default()
	productsInfra.RegisterProductRoutes(router)
	router.Run(":8080")
}