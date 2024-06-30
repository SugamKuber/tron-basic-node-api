package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"tron-api-server/internal/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env")
		return
	}

	router := gin.Default()

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to run server on 8080")
	}
}
