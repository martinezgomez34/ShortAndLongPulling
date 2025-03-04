package infrastructure

import (
	"github.com/gin-gonic/gin"
	MySQL "github.com/martinezgomez34/abarrote/src/core"
	"github.com/martinezgomez34/abarrote/src/products/application"
	controllers "github.com/martinezgomez34/abarrote/src/products/infraestructure/controller"
	"github.com/martinezgomez34/abarrote/src/products/infraestructure/rute"
)

func RegisterProductRoutes(router *gin.Engine) {
	db := MySQL.GetDBPool()


	repo := NewSQL(db)
	newProductUseCase := application.NewProductUC(repo)
	listProductUseCase := application.ListProductUC(repo)
	getByIDProductUseCase := application.GetByIDProductUC(repo)
	editProductUseCase := application.EditProductUC(repo)
	deleteProductUseCase := application.DeleteProductUC(repo)

	notifyProductCreatedUseCase := application.NewNotifyProductCreatedUC(repo)
	notifyProductActionUseCase := application.NewNotifyProductActionUC(repo)
	notifyLowStockUseCase := application.NewNotifyLowStockUC(repo)

	newProductController := controllers.NewProductController(newProductUseCase)
	listProductController := controllers.NewListProductController(listProductUseCase)
	getByIDProductController := controllers.NewGetByIDProductController(getByIDProductUseCase)
	editProductController := controllers.NewEditProductController(editProductUseCase)
	deleteProductController := controllers.NewDeleteProductController(deleteProductUseCase)

	notifyProductCreatedController := controllers.NewNotifyProductCreatedController(notifyProductCreatedUseCase)
	notifyProductActionController := controllers.NewNotifyProductActionController(notifyProductActionUseCase)
	notifyLowStockController := controllers.NewNotifyLowStockController(notifyLowStockUseCase)

	rute.ProductRoutes(router, newProductController, listProductController, editProductController, getByIDProductController, 
		deleteProductController, notifyProductCreatedController, notifyProductActionController, notifyLowStockController)
}
