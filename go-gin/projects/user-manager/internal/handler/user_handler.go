package handler

import (
	"github.com/danielleit241/internal/dto"
	"github.com/danielleit241/internal/service"
	"github.com/danielleit241/internal/utils"
	"github.com/danielleit241/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService service.UserService
}

type GetUserByIDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type GetAllUsersRequest struct {
	Query string `form:"query" binding:"omitempty,search"`
	Page  *int   `form:"page" binding:"omitempty,gt=0"`
	Limit *int   `form:"limit" binding:"omitempty,gt=0,lte=100"`
}

func NewUserHandler(userService service.UserService) *UserHandler {
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

	page, limit := utils.NormalizePaginationParams(req.Page, req.Limit)

	users, total, err := uh.userService.GetAllUsersWithPagination(req.Query, page, limit)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	responses := dto.ToResponses(users)

	utils.ResponseSuccessWithPage(c, "users retrieved successfully", total, page, limit, responses)
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

	utils.ResponseSuccess(c, 200, response, "user retrieved successfully")
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user dto.UserCreateRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ResponseValidationError(c, validation.HandleValidationError(err))
		return
	}

	entity := user.ToEntity(user)

	createdUser, err := uh.userService.CreateUser(entity)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	response := dto.ToResponse(createdUser)

	utils.ResponseSuccess(c, 201, response, "user created successfully")
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
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

	var user dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ResponseValidationError(c, validation.HandleValidationError(err))
		return
	}

	entity := user.ToEntity(user)

	updated, err := uh.userService.UpdateUser(userID, entity)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	response := dto.ToResponse(updated)

	utils.ResponseSuccess(c, 200, response, "user updated successfully")
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
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

	err = uh.userService.DeleteUser(userID)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, 204, nil, "user deleted successfully")
}
