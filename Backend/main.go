package main

import (
	"Crud_fiber_Go/config"
	"Crud_fiber_Go/models"
	"Crud_fiber_Go/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	// Connect to database
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		log.Fatal("DB_PASS environment variable is not set")
	}
	// Load configuration using Builder Pattern
	appConfig := config.NewAppConfigBuilder().
		SetPort(":3001").
		SetDatabase(dbPass).
		Build()

	// Initialize Fiber App
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Allow your frontend origin
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))
	models.ConnectDatabase()
	// Connect to database
	if err := appConfig.ConnectDB(); err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Register routes
	routes.RegisterRoutes(app)

	// Start server
	log.Fatal(app.Listen(appConfig.Port))
}
