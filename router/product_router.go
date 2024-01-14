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

	// Main product router group
	productRouter := app.Group("/product")

	// Normal product, without auth
	productRouter.Get("", productCtrl.FindAll)
	// productRouter.Get("/:id", productCtrl.FindById)
	// productRouter.Post("/", productCtrl.Create)
	// productRouter.Put("/:id", productCtrl.Update)
	// productRouter.Delete("/:id", productCtrl.Delete)
}
