package repository

import (
	"github.com/danielleit241/internal/models"
	"github.com/danielleit241/internal/utils"
	"github.com/google/uuid"
)

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

func (repo *inMemoryUserRepository) FindAllWithPagination(query string, page, limit int) ([]models.User, int, error) {

	var filteredUsers []models.User

	if query == "" {
		filteredUsers = repo.users
	} else {
		for _, user := range repo.users {
			if utils.ContainsIgnoreCase(user.Name, query) ||
				utils.ContainsIgnoreCase(user.Email, query) {

				filteredUsers = append(filteredUsers, user)
			}
		}
	}

	total := len(filteredUsers)

	start := (page - 1) * limit
	if start >= total {
		return []models.User{}, total, nil
	}

	end := min(start+limit, total)

	return filteredUsers[start:end], total, nil
}

func (repo *inMemoryUserRepository) FindById(id uuid.UUID) (*models.User, error) {
	for _, user := range repo.users {
		if user.ID == id {
			return &user, nil
		}
	}
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
