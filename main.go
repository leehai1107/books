package main

import (
	"example/Demo/initilization"
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

	db, redisClient, err := initilization.ConnectToDatabase()
	if err != nil {
		log.Fatal("Error connecting to the database and/or Redis:", err)
	}

	// Close Redis client when the main function exits
	defer redisClient.Close()

	// Create a new Gin router
	router := routes.CreateRoutes(db)

	router.Run("localhost:8080")
}
