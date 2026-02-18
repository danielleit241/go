package repository

import "github.com/danielleit241/entity"

type UserRepository interface {
	Repository[entity.User]
	FindUserByEmail(email string) ([]entity.User, error)
}

type inMemoryUserRepository struct {
	*inMemoryRepository[entity.User]
}

func NewInMemoryUserRepository() *inMemoryUserRepository {
	return &inMemoryUserRepository{
		inMemoryRepository: NewInMemoryRepository[entity.User](),
	}
}

func (r *inMemoryUserRepository) FindUserByEmail(email string) ([]entity.User, error) {
	var results []entity.User
	for _, user := range r.items {
		if user.Email == email {
			results = append(results, user)
		}
	}
	return results, nil
}
