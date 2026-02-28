package handler

import (
	"net/http"

	"github.com/danielleit241/internal/dto"
	"github.com/danielleit241/internal/models"
	"github.com/danielleit241/internal/service"
	"github.com/danielleit241/internal/utils"
	"github.com/danielleit241/internal/validation"
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
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, validation.HandleValidationError(err))
		return
	}

	createdUser, err := uh.userService.CreateUser(user)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	response := dto.MapUserToResponse(createdUser)

	utils.ResponseSuccess(c, response, "user created successfully")
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
}
