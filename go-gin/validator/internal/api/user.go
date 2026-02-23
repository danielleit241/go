package api

import (
	"net/http"
	"os"

	"github.com/danielleit241/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

type GetByUserIDParam struct {
	ID int `uri:"id" binding:"gt=0"`
}

type CreateUserRequest struct {
	Username string `form:"username" binding:"required,min=3,max=50"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=8"`
}

type UpdateAvatarRequest struct {
	AvatarName string `form:"avatar_name" binding:"required"`
}

func NewUserController() *UserController {
	return &UserController{}
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all users - v1",
	})
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
	var param GetByUserIDParam
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	c.JSON(200, gin.H{
		"message": "Get user by ID - v1",
		"id":      param.ID,
	})
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	c.JSON(200, gin.H{
		"message": "Create user - v1",
		"data":    req,
	})
}

func (ctrl *UserController) UpdateUserAvatar(c *gin.Context) {
	var param GetByUserIDParam
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	var req UpdateAvatarRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	image, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Avatar image is required"})
		return
	}

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create uploads directory"})
		return
	}

	dst := "./uploads/" + image.Filename
	if err := c.SaveUploadedFile(image, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Update user avatar - v1",
		"id":          param.ID,
		"avatar_name": req.AvatarName,
		"avatar_file": image.Filename,
	})
}
