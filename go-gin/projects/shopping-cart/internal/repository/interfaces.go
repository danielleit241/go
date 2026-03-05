package repository

import (
	"context"

	"github.com/danielleit241/internal/db/sqlc"
)

type UserRepository interface {
	FindAll()
	FindAllWithPagination(query string, page, limit int)
	FindById()
	Create(ctx context.Context, userCreateParams sqlc.CreateUserParams) (sqlc.User, error)
	Update()
	Delete()
	IsEmailExists(email string)
}
