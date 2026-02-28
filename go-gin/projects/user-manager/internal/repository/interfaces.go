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
	Update(id uuid.UUID, user models.User) (*models.User, error)
	Delete(id uuid.UUID) error
	IsEmailExists(email string) (*models.User, bool)
}
