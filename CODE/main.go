package main

import (
	"dev/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Third experiment
func main() {
	app := fiber.New()

	// Apply CORS middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Method() == "OPTIONS" {
			return c.Status(fiber.StatusNoContent).SendString("")
		}

		return c.Next()
	})

	// Serve static files from the "public" directory
	app.Static("/public", "./public")

	router.ProductRouter(app)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
