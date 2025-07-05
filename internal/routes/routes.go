package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/user-management-api/internal/middleware"
)

// Route is an interface for registering grouped routes like /users, /products, etc.
type Route interface {
	Register(r *gin.RouterGroup)
}

// RegisterRoutes applies all route modules to the main router group
func RegisterRoutes(r *gin.Engine, routeModules ...Route) {
	api := r.Group("/api/v1") // You can make this configurable
	api.Use(middleware.AuthMiddleware("your-secret-key"))

	for _, route := range routeModules {
		route.Register(api)
	}
}
