package repository

import "github.com/danielleit241/internal/models"

type userRepository struct {
	users []models.User
}

func NewUserRepository() UserRepository {
	return &userRepository{
		users: []models.User{},
	}
}

func (repo *userRepository) GetUserByID() {
}

func (repo *userRepository) AddUser() {
}
