package main

import (
	"log"

	"github.com/bchaillou003/marvel-family-backend/database"
	"github.com/bchaillou003/marvel-family-backend/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Loading .env file to load environment variables
	godotenv.Load()

	// Connection to database
	database.ConnectDb()

	// Initialize app fiber-
	app := fiber.New()
	app.Use(cors.New())

	// Initialize routes
	server.SetupRoutes(app)

	// Listen serves HTTP requests from the given addr.
	log.Fatal(app.Listen(":3000"))
}
