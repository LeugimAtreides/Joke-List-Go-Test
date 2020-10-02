package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// function to retrieve env variables
func goDotEnvVariable(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	// Set the router as the default shipped with Gin
	router := gin.Default()

	router.Run(":5000")
}
