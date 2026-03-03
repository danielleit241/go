package repository

import (
	"fmt"

	"github.com/danielleit241/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return &user, nil
}

func (repo *userRepository) AddUser(user *models.User) error {
	var existingUser models.User
	if err := repo.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	} else if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to check existing user: %w", err)
	}
	return repo.db.Create(user).Error
}
