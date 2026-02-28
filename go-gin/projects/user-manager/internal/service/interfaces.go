package service

import (
	"github.com/danielleit241/internal/models"
	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	CreateUser(user models.User) (*models.User, error)
	UpdateUser(id int, user models.User) (*models.User, error)
	DeleteUser(id int) error
}
