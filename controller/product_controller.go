package controller

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	FindAll(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	// Update(ctx *gin.Context)
	// Delete(ctx *gin.Context)
}
