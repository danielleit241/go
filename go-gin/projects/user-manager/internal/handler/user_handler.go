package handler

import (
	"github.com/danielleit241/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	uh.userService.GetAllUsers()
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
}
