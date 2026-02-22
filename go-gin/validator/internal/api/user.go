package api

import (
	"net/http"

	"github.com/danielleit241/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

type GetByUserIDParam struct {
	ID int `uri:"id" binding:"gt=0"`
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
	c.JSON(200, gin.H{
		"message": "Create user - v1",
	})
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update user - v1",
	})
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete user - v1",
	})
}
