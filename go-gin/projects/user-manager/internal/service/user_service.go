package service

import "github.com/danielleit241/internal/repository"

type UserService struct {
	inMemoryRepo *repository.InMemoryUserRepository
}

func NewUserService(inMemoryRepo *repository.InMemoryUserRepository) *UserService {
	return &UserService{
		inMemoryRepo: inMemoryRepo,
	}
}
