package service

import "github.com/danielleit241/internal/models"

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(user models.User) (*models.User, error)
	UpdateUser(id int, user models.User) (*models.User, error)
	DeleteUser(id int) error
}
