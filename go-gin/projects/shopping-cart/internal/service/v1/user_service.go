package v1service

import (
	"github.com/danielleit241/internal/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetAllUsersWithPagination(query string, page, limit int) {
}

func (us *userService) GetAllUsers() {
}

func (us *userService) GetUserByID() {
}

func (us *userService) CreateUser() {
}

func (us *userService) UpdateUser() {
}

func (us *userService) DeleteUser() {
}
