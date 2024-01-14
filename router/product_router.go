package router

import (
	"dev/app"
	"dev/controller"
	"dev/repository"

	"github.com/gin-gonic/gin"
)

func ProductRouter(router *gin.Engine) {
	db := app.NewDB()

	productRepo := repository.NewProductRepository(db)

	productCtrl := controller.NewProductController(productRepo)

	// Main product router group
	productRouter := router.Group("/product")

	// Normal product, without auth
	productRouter.GET("", productCtrl.FindAll)
	productRouter.GET("/:id", productCtrl.FindById)
	productRouter.POST("/", productCtrl.Create)
	productRouter.PUT("/:id", productCtrl.Update)
}
