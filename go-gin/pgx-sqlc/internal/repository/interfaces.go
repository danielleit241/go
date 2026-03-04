package repository

import (
	"context"

	"github.com/danielleit241/internal/db/sqlc"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*sqlc.User, error)
	AddUser(ctx context.Context, user sqlc.CreateUserParams) (*sqlc.User, error)
}
