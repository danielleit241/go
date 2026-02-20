package v1handler

import (
	"strconv"

	"github.com/danielleit241/utils"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "Alice", Age: 30},
	{ID: 2, Name: "Bob", Age: 25},
	{ID: 3, Name: "Charlie", Age: 35},
}

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUsersV1(c *gin.Context) {
	c.JSON(200, gin.H{
		"users": users,
	})
}

func (h *UserHandler) GetUserByIDV1(c *gin.Context) {
	idParam := c.Param("id")

	id, err := utils.ValidationPositiveInteger("id", idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(200, gin.H{
				"user": user,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error": "User not found",
	})
}

func (h *UserHandler) CreateUserV1(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	c.JSON(201, gin.H{
		"user": newUser,
	})
}

func (h *UserHandler) UpdateUserV1(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			c.JSON(200, gin.H{
				"user": updatedUser,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error": "User not found",
	})
}

func (h *UserHandler) DeleteUserV1(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(200, gin.H{
				"message": "User deleted",
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error": "User not found",
	})
}
