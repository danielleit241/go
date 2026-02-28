package handler

import (
	"github.com/danielleit241/internal/dto"
	"github.com/danielleit241/internal/models"
	"github.com/danielleit241/internal/service"
	"github.com/danielleit241/internal/utils"
	"github.com/danielleit241/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	users, err := uh.userService.GetAllUsers()
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	responses := dto.ToResponses(users)

	utils.ResponseSuccess(c, responses, "users retrieved successfully")
}

type GetUserByIDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	var param GetUserByIDRequest
	if err := c.ShouldBindUri(&param); err != nil {
		utils.ResponseValidationError(c, validation.HandleValidationError(err))
		return
	}

	userID, err := uuid.Parse(param.ID)
	if err != nil {
		utils.ResponseValidationError(c, validation.HandleValidationError(err))
		return
	}

	user, err := uh.userService.GetUserByID(userID)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	response := dto.ToResponse(user)

	utils.ResponseSuccess(c, response, "user retrieved successfully")
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ResponseValidationError(c, validation.HandleValidationError(err))
		return
	}

	createdUser, err := uh.userService.CreateUser(user)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	response := dto.ToResponse(createdUser)

	utils.ResponseSuccess(c, response, "user created successfully")
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
}
