package controller

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	FindAll(ctx *fiber.Ctx) error
	// FindById(ctx *gin.Context)
	// Create(ctx *gin.Context)
	// Update(ctx *gin.Context)
	// Delete(ctx *gin.Context)
}
