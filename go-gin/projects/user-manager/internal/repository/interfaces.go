package repository

import "github.com/danielleit241/internal/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindById(id int) (*models.User, error)
	Create(user models.User) (*models.User, error)
	Update(id int, user models.User) (*models.User, error)
	Delete(id int) error
	IsEmailExists(email string) bool
}
