package main

import (
	"example/Demo/initializers"
	"example/Demo/routes"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := initializers.ConnectToDatabase()

	// Create a new Gin router
	router := routes.CreateRoutes(db)

	router.Run("localhost:8080")
}
