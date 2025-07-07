package repository

import "github.com/thinhcompany/user-management-api/internal/models"

type IUserRepository interface {
	Create(user *models.User) *models.User
	GetByID(id int) (*models.User, bool)
	GetAll() []*models.User
}
