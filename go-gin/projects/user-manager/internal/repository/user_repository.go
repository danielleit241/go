package repository

import "github.com/danielleit241/internal/models"

type inMemoryUserRepository struct {
	users []models.User
}

func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{
		users: make([]models.User, 0),
	}
}

func (repo *inMemoryUserRepository) FindAll() ([]models.User, error) {
	return repo.users, nil
}

func (repo *inMemoryUserRepository) FindById(id int) (*models.User, error) {
	return nil, nil
}

func (repo *inMemoryUserRepository) Create(user models.User) (*models.User, error) {
	repo.users = append(repo.users, user)
	return &user, nil
}

func (repo *inMemoryUserRepository) Update(id int, user models.User) (*models.User, error) {
	return nil, nil
}

func (repo *inMemoryUserRepository) Delete(id int) error {
	return nil
}

func (repo *inMemoryUserRepository) IsEmailExists(email string) bool {
	for _, user := range repo.users {
		if user.Email == email {
			return true
		}
	}
	return false
}
