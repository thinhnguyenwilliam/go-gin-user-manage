package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thinhcompany/user-management-api/internal/config"
	"github.com/thinhcompany/user-management-api/internal/handler"
	"github.com/thinhcompany/user-management-api/internal/repository"
	"github.com/thinhcompany/user-management-api/internal/routes"
	"github.com/thinhcompany/user-management-api/internal/service"
)

func main() {
	// Load .env file (optional)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration
	cfg := config.NewConfig()

	// Initialize dependencies
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userHandler)

	// Set up Gin router
	r := gin.Default()

	// Register all route modules
	routes.RegisterRoutes(r, userRoutes)

	// Start server
	fmt.Printf("Server running on port %s in %s mode\n", cfg.AppPort, cfg.AppEnv)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
