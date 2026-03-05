package v1service

import (
	"errors"

	"github.com/danielleit241/internal/db/sqlc"
	"github.com/danielleit241/internal/repository"
	"github.com/danielleit241/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetAllUsersWithPagination(query string, page, limit int) {
}

func (us *userService) GetAllUsers() {
}

func (us *userService) GetUserByID() {
}

func (us *userService) CreateUser(ctx *gin.Context, userCreateParams sqlc.CreateUserParams) (sqlc.User, error) {
	context := ctx.Request.Context()

	userCreateParams.Email = utils.NormalizeString(userCreateParams.Email)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userCreateParams.Password), bcrypt.DefaultCost)
	if err != nil {
		return sqlc.User{}, utils.WrapError("failed to hash password", utils.ErrCodeInternalServerError, err)
	}
	userCreateParams.Password = string(hashedPassword)

	createdUser, err := us.repo.Create(context, userCreateParams)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return sqlc.User{}, utils.WrapError("email already exists", utils.ErrCodeConflict, err)
		}
		return sqlc.User{}, utils.WrapError("failed to create user", utils.ErrCodeInternalServerError, err)
	}

	return createdUser, nil
}

func (us *userService) UpdateUser() {
}

func (us *userService) DeleteUser() {
}
