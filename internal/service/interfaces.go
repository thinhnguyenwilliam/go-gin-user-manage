package service

import "github.com/thinhcompany/user-management-api/internal/models"

// User service interface for abstraction
type IUserService interface {
	CreateUser(name, email, password string) *models.User
	GetUserByID(id int) (*models.User, bool)
	GetAllUsers() []*models.User
}
