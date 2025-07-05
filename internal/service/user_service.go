package service

import (
	"github.com/thinhcompany/user-management-api/internal/models"
	"github.com/thinhcompany/user-management-api/internal/repository"
)

type UserService struct {
	repo *repository.InMemoryUserRepository
}

func NewUserService(repo *repository.InMemoryUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(name, email, password string) *models.User {
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: password, // In real apps, hash this!
	}
	return s.repo.Create(user)
}

func (s *UserService) GetUserByID(id int) (*models.User, bool) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetAllUsers() []*models.User {
	return s.repo.GetAll()
}
