package v2handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUsersV2(c *gin.Context) {
	c.JSON(200, gin.H{
		"users": "list of users (v2)",
	})
}

func (h *UserHandler) GetUserByIDV2(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": "user details (v2)",
	})
}

func (h *UserHandler) CreateUserV2(c *gin.Context) {
	c.JSON(201, gin.H{
		"user": "new user created (v2)",
	})
}

func (h *UserHandler) UpdateUserV2(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": "user updated (v2)",
	})
}

func (h *UserHandler) DeleteUserV2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user deleted (v2)",
	})
}
