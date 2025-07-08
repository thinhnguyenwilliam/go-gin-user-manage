package app

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/user-management-api/internal/config"
	"github.com/thinhcompany/user-management-api/internal/routes"
)

type Application struct {
	Config *config.Config
	Router *gin.Engine
}

func NewApplication() *Application {
	cfg := config.NewConfig()
	r := gin.Default()

	// Collect all modules that implement Route interface
	modules := []routes.Route{
		NewUserModule(),
		// Add more modules here like: NewProductModule(), NewCategoryModule()
	}

	// Register all module routes
	routes.RegisterRoutes(r, modules...)

	return &Application{
		Config: cfg,
		Router: r,
	}
}

func (a *Application) Run() error {
	address := ":" + a.Config.AppPort
	return a.Router.Run(address)
}
