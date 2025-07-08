package app

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/user-management-api/internal/handler"
	"github.com/thinhcompany/user-management-api/internal/repository"
	"github.com/thinhcompany/user-management-api/internal/routes"
	"github.com/thinhcompany/user-management-api/internal/service"
)

type UserModule struct {
	repo    repository.IUserRepository
	service service.IUserService
	handler *handler.UserHandler // ✅ use pointer
}

// Ensure UserModule implements routes.Route
var _ routes.Route = (*UserModule)(nil)

func NewUserModule() *UserModule {
	repo := repository.NewUserRepository()
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	return &UserModule{
		repo:    repo,
		service: service,
		handler: handler,
	}
}

// Register satisfies the routes.Route interface
func (m *UserModule) Register(r *gin.RouterGroup) {
	routes.RegisterUserRoutes(r, m.handler)
}
