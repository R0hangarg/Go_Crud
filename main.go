package main

import (
	"Crud_fiber_Go/config"
	"Crud_fiber_Go/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		log.Fatal("DB_PASS environment variable is not set")
	}
	// Load configuration using Builder Pattern
	appConfig := config.NewAppConfigBuilder().
		SetPort(":3000").
		SetDatabase(dbPass).
		Build()

	// Initialize Fiber App
	app := fiber.New()

	// Connect to database
	if err := appConfig.ConnectDB(); err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Register routes
	routes.RegisterRoutes(app)

	// Start server
	log.Fatal(app.Listen(appConfig.Port))
}
