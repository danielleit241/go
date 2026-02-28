package repository

import (
	"slices"
	"strings"

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
	for i := range repo.users {
		if repo.users[i].ID == id {
			return &repo.users[i], nil
		}
	}
	return nil, utils.WrapError("failed to find user by id", utils.ErrCodeNotFound, nil)
}

func (repo *inMemoryUserRepository) Create(user models.User) (*models.User, error) {
	repo.users = append(repo.users, user)
	return &user, nil
}

func (repo *inMemoryUserRepository) Update(id uuid.UUID, user models.User) (*models.User, error) {
	updated, err := repo.FindById(id)
	if err != nil {
		return nil, err
	}

	if user.Name != "" {
		updated.Name = strings.TrimSpace(user.Name)
	}
	if user.Email != "" {
		updated.Email = utils.NormalizeString(user.Email)
	}
	if user.Age != 0 {
		updated.Age = user.Age
	}
	if user.Status != 0 {
		updated.Status = user.Status
	}
	if user.Level != 0 {
		updated.Level = user.Level
	}

	return updated, nil
}

func (repo *inMemoryUserRepository) Delete(id uuid.UUID) error {
	for i := range repo.users {
		if repo.users[i].ID == id {
			repo.users = slices.Delete(repo.users, i, i+1)
			return nil
		}
	}
	return utils.WrapError("failed to delete user", utils.ErrCodeNotFound, nil)
}

func (repo *inMemoryUserRepository) IsEmailExists(email string) (*models.User, bool) {
	for _, user := range repo.users {
		if user.Email == email {
			return &user, true
		}
	}
	return nil, false
}
