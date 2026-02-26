package repository

import "github.com/danielleit241/internal/models"

type InMemoryUserRepository struct {
	users []models.User
}

func NewUserInMemoryRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make([]models.User, 0),
	}
}
