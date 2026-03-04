package handler

import (
	"net/http"

	"github.com/danielleit241/internal/repository"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, "User retrieved successfully")
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, "User created successfully")
}
