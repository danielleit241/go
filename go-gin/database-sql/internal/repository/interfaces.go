package repository

import "github.com/danielleit241/internal/models"

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
	AddUser(user *models.User) error
}
