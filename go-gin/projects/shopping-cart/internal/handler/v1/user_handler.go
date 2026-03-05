package v1handler

import (
	v1dto "github.com/danielleit241/internal/dto/v1"
	v1service "github.com/danielleit241/internal/service/v1"
	"github.com/danielleit241/internal/utils"
	"github.com/danielleit241/internal/validation"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService v1service.UserService
}

type GetUserByIDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type GetAllUsersRequest struct {
	Query string `form:"query" binding:"omitempty,search"`
	Page  *int   `form:"page" binding:"omitempty,gt=0"`
	Limit *int   `form:"limit" binding:"omitempty,gt=0,lte=100"`
}

func NewUserHandler(userService v1service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	var req GetAllUsersRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseValidationError(c, validation.HandleValidationError(err))
		return
	}

	// page, limit := utils.NormalizePaginationParams(req.Page, req.Limit)

	// utils.ResponseSuccessWithPage(c, "users retrieved successfully", total, page, limit, responses)
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	var param GetUserByIDRequest
	if err := c.ShouldBindUri(&param); err != nil {
		utils.ResponseValidationError(c, validation.HandleValidationError(err))
		return
	}

	utils.ResponseSuccess(c, 200, nil, "user retrieved successfully")
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user v1dto.UserCreateRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ResponseValidationError(c, validation.HandleValidationError(err))
		return
	}

	createdUser, err := uh.userService.CreateUser(c, user.ToCreateEntity())
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	response := v1dto.ToResponse(createdUser)

	utils.ResponseSuccess(c, 201, response, "user created successfully")
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var param GetUserByIDRequest
	if err := c.ShouldBindUri(&param); err != nil {
		utils.ResponseValidationError(c, validation.HandleValidationError(err))
		return
	}

	var user v1dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ResponseValidationError(c, validation.HandleValidationError(err))
		return
	}

	utils.ResponseSuccess(c, 200, nil, "user updated successfully")
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	var param GetUserByIDRequest
	if err := c.ShouldBindUri(&param); err != nil {
		utils.ResponseValidationError(c, validation.HandleValidationError(err))
		return
	}

	utils.ResponseSuccess(c, 204, nil, "user deleted successfully")
}
