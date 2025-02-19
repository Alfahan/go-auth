package main

import (
	"go-auth/configs"
	"go-auth/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	// Load .env file and Create a new connection to the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := configs.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Schema:   os.Getenv("SCHEMA"),
	}

	// initialize DB
	configs.InitDB(config)

	// load the routes
	routes.AuthRoutes(r)

	// run the server
	r.Run(":8080")
}
