package repository

import (
	"database/sql"

	"github.com/danielleit241/internal/models"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (repo *userRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}

	err := repo.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *userRepository) AddUser(user *models.User) error {
	row := repo.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email)

	err := row.Scan(&user.ID)

	if err != nil {
		panic(err)
	}

	return nil
}
