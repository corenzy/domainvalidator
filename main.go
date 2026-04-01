package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/corenzy/domainvalidator/handlers"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Domain Validator API",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":    "Domain Validator API",
			"version": "1.0.0",
			"lib":  "com.corenzy.domainvalidator"
		})
	})

	app.Post("/lookup", handlers.HandleLookup)

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Domain Validator API running on :%s", port)
	log.Fatal(app.Listen(":" + port))
}
