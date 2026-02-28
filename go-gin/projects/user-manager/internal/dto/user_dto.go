package dto

import (
	"github.com/danielleit241/internal/models"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Age    int       `json:"age"`
	Status string    `json:"status"`
	Level  string    `json:"level"`
}

func MapUserToResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
		Status: getStringStatus(user.Status),
		Level:  getStringLevel(user.Level),
	}
}

func getStringStatus(status int) string {
	if status == 1 {
		return "active"
	}
	return "inactive"
}

func getStringLevel(level int) string {
	if level == 2 {
		return "admin"
	}
	return "user"
}
