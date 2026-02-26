package service

import (
	"github.com/danielleit241/internal/models"
	"github.com/danielleit241/internal/repository"
)

type userService struct {
	inMemoryRepo repository.UserRepository
}

func NewUserService(inMemoryRepo repository.UserRepository) UserService {
	return &userService{
		inMemoryRepo: inMemoryRepo,
	}
}

func (us *userService) GetAllUsers() ([]models.User, error) {
	return []models.User{}, nil
}

func (us *userService) GetUserByID(id int) (*models.User, error) {
	return nil, nil
}

func (us *userService) CreateUser(user models.User) (*models.User, error) {
	return nil, nil
}

func (us *userService) UpdateUser(id int, user models.User) (*models.User, error) {
	return nil, nil
}

func (us *userService) DeleteUser(id int) error {
	return nil
}
