package router

import (
	"dev/app/config"
	"dev/controller"
	"dev/repository"

	"github.com/gofiber/fiber/v2"
)

func ProductRouter(app *fiber.App) {
	db := config.NewDB()

	productRepo := repository.NewProductRepository(db)

	productCtrl := controller.NewProductController(productRepo)

	productRouter := app.Group("/product")

	productRouter.Get("", productCtrl.FindAll)
	productRouter.Get("/:id", productCtrl.FindById)
	productRouter.Post("/", productCtrl.Create)
	productRouter.Put("/:id", productCtrl.Update)
	productRouter.Delete("/:id", productCtrl.Delete)

	// Example usage for paging
	// http://localhost:8080/product?page=2&pageSize=4
	// use it in here productRouter.Get("", productCtrl.FindAll)
}
