package repository

import (
	"github.com/danielleit241/internal/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindAllWithPagination(query string, page, limit int) ([]models.User, int, error)
	FindById(id uuid.UUID) (*models.User, error)
	Create(user models.User) (*models.User, error)
	Update(id int, user models.User) (*models.User, error)
	Delete(id int) error
	IsEmailExists(email string) bool
}
