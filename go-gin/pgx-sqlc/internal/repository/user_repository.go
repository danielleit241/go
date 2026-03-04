package repository

import (
	"context"

	"github.com/danielleit241/internal/db/sqlc"
	"github.com/google/uuid"
)

type userRepository struct {
	db sqlc.Querier
}

func NewUserRepository(db sqlc.Querier) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*sqlc.User, error) {
	user, err := repo.db.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) AddUser(ctx context.Context, user sqlc.CreateUserParams) (*sqlc.User, error) {
	userCreated, err := repo.db.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &userCreated, nil
}
