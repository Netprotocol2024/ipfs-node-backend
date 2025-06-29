package main

import (
	"log"
	"os"

	"vps-monitor/controllers"
	"vps-monitor/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName:               "VPS Monitor API v1.0",
		DisableStartupMessage: true,
	})

	// Middleware
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path}\n",
	}))
	app.Use(recover.New())

	// Allow all CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// Dependency Injection
	systemService := services.NewSystemService()
	metricsController := controllers.NewMetricsController(systemService)

	// API routes with versioning
	api := app.Group("/api/v1")

	// Metrics endpoints
	api.Get("/metrics/vps", metricsController.GetMetricsInfo)
	api.Get("/metrics/storage", metricsController.GetStorageInfo)

	// Health check endpoint
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": "Online",
		})
	})

	// Not found handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Endpoint not found",
		})
	})

	// Get port from environment or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Graceful startup message
	log.Printf("🚀 Server started on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
