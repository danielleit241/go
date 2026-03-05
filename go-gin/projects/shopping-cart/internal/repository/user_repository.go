package repository

import (
	"context"

	"github.com/danielleit241/internal/db/sqlc"
)

type userRepository struct {
	db sqlc.Querier
}

func NewUserRepository(db sqlc.Querier) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) FindAll() {

}

func (repo *userRepository) FindAllWithPagination(query string, page, limit int) {
}

func (repo *userRepository) FindById() {

}

func (repo *userRepository) Create(ctx context.Context, userCreateParams sqlc.CreateUserParams) (sqlc.User, error) {
	created, err := repo.db.CreateUser(ctx, userCreateParams)
	if err != nil {
		return sqlc.User{}, err
	}

	return created, nil
}

func (repo *userRepository) Update() {

}

func (repo *userRepository) Delete() {
}

func (repo *userRepository) IsEmailExists(email string) {

}
