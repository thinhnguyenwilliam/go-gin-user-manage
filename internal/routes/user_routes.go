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

func (u *UserRoutes) Register(r *gin.RouterGroup) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", u.handler.CreateUser)
		userGroup.GET("/", u.handler.GetAllUsers)
		userGroup.GET("/:uuid", u.handler.GetUserByID)
	}
}
