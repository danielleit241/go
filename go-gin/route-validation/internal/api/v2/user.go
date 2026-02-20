package v2handler

import (
	"net/http"

	"github.com/danielleit241/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserV2 struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

var usersV2 = []UserV2{
	{ID: uuid.New(), Name: "Alice", Age: 30},
	{ID: uuid.New(), Name: "Bob", Age: 25},
	{ID: uuid.New(), Name: "Charlie", Age: 35},
}

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUsersV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": usersV2,
	})
}

func (h *UserHandler) GetUserByIDV2(c *gin.Context) {
	uuidParam := c.Param("uuid")

	// Validate that the ID is a valid UUID
	u, err := utils.ValidationUUID("id", uuidParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	for _, user := range usersV2 {
		if user.ID == u {
			c.JSON(http.StatusOK, gin.H{
				"user": user,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "User not found",
	})
}

func (h *UserHandler) CreateUserV2(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"user": "new user created (v2)",
	})
}

func (h *UserHandler) UpdateUserV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "user updated (v2)",
	})
}

func (h *UserHandler) DeleteUserV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted (v2)",
	})
}
