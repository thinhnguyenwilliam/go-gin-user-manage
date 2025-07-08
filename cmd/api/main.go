package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/thinhcompany/user-management-api/internal/app"
)

func main() {
	// Load .env file (optional)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	application := app.NewApplication()

	fmt.Printf("Server running on port %s in %s mode\n", application.Config.AppPort, application.Config.AppEnv)

	// Start server
	if err := application.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
