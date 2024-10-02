package main

import (
	"fmt"
	"github.com/gatij/gin-golang-app/internal/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	// Initialize a new Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server on port 8000
	err := router.Run(":8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
