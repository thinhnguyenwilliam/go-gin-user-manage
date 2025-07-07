package service

import (
	"log"

	"github.com/thinhcompany/user-management-api/internal/models"
	"github.com/thinhcompany/user-management-api/internal/repository"
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
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
	log.Println("User Service is here: GetAllUsers")
	return s.repo.GetAll()
}
