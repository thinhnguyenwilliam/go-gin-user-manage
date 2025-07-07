package repository

import (
	"log"

	"github.com/thinhcompany/user-management-api/internal/models"
)

type InMemoryUserRepository struct {
	users  map[int]*models.User
	autoID int
}

func NewUserRepository() IUserRepository {
	return &InMemoryUserRepository{
		users:  make(map[int]*models.User),
		autoID: 1,
	}
}

func (repo *InMemoryUserRepository) Create(user *models.User) *models.User {
	user.ID = repo.autoID
	repo.users[user.ID] = user
	repo.autoID++
	return user
}

func (repo *InMemoryUserRepository) GetByID(id int) (*models.User, bool) {
	user, exists := repo.users[id]
	return user, exists
}

func (repo *InMemoryUserRepository) GetAll() []*models.User {
	log.Println("User Repository is here: GetAllUsers")
	users := []*models.User{}
	for _, user := range repo.users {
		users = append(users, user)
	}
	return users
}
