package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default shipped with Gin
	router := gin.Default()

	router.Run(":5000")
}
