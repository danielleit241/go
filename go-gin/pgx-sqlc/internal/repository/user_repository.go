package repository

import (
	"github.com/danielleit241/internal/db/sqlc"
)

type userRepository struct {
	db *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) GetUserByID() {
}

func (repo *userRepository) AddUser() {
}
