package main

import (
	"log"
	"myapi/configs"
	"myapi/middlewares"

	"myapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	configs.LoadConfig()

	// Connect to the database
	configs.ConnectDatabase()

	// Set Gin mode
	mode := configs.GetGinMode()
	gin.SetMode(mode)

	// Create a new Gin instance
	r := gin.New()

	// Middlewares
	r.Use(middlewares.CustomLogger())
	r.Use(gin.Recovery())

	// Routes
	routes.SetupRootRouter(r)

	// Run the server
	port := configs.GetPort()
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
