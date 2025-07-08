package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/user-management-api/internal/handler"
)

type UserRoutes struct {
	handler *handler.UserHandler
}

func NewUserRoutes(handler *handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		handler: handler,
	}
}

func RegisterUserRoutes(r *gin.RouterGroup, h *handler.UserHandler) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", h.CreateUser)
		userGroup.GET("/", h.GetAllUsers)
		userGroup.GET("/:uuid", h.GetUserByID)
	}
}
