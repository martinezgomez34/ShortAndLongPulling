package rute

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/martinezgomez34/abarrote/src/products/infraestructure/controller"
)

func ProductRoutes(router *gin.Engine, ctrlCreate *controllers.CreateProductController, ctrlGetAll *controllers.ListProductController,  
	ctrlEdit *controllers.EditProductController, ctrlGetByID *controllers.GetByIDProductController, ctrlDelete *controllers.DeleteProductController,
	ctrlNotifyCreate *controllers.NotifyProductCreatedController, ctrlNotifyAction *controllers.NotifyProductActionController, ctrlNotifyLowStock *controllers.NotifyLowStockController) {
	group := router.Group("/products")


	{
		group.POST("/", ctrlCreate.CreateProduct)
		group.GET("/", ctrlGetAll.GetAllProducts)
		group.GET("/:id", ctrlGetByID.GetProductByID)
		group.PUT("/:id", ctrlEdit.UpdateProduct)
		group.DELETE("/:id", ctrlDelete.DeleteProduct)
		group.GET("/notify/new", ctrlNotifyCreate.WaitForNewProduct)
		group.GET("/notify/action", ctrlNotifyAction.WaitForAction)
		group.GET("/notify/low-stock", ctrlNotifyLowStock.GetLowStockProducts)
	}
}
