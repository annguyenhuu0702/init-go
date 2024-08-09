package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// get env file
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// get port
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // giá trị mặc định
	}
	return port
}

// set gin mode
func GetGinMode() string {
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug" // giá trị mặc định
	}
	return mode
}
