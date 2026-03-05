package v1service

import (
	"github.com/danielleit241/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetAllUsers()
	GetAllUsersWithPagination(query string, page, limit int)
	GetUserByID()
	CreateUser(ctx *gin.Context, userCreateParams sqlc.CreateUserParams) (sqlc.User, error)
	UpdateUser()
	DeleteUser()
}
