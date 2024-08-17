package main

import (
	"github.com/bchaillou003/marvel-family-backend/server"
	"github.com/joho/godotenv"
)

func main() {
	// Loading .env file to load environment variables
	godotenv.Load()

	api := server.NewAPI()

	defer api.App.Close()

	api.Run()
}
