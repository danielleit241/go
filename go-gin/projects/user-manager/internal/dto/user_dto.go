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

func ToResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
		Status: getStringStatus(user.Status),
		Level:  getStringLevel(user.Level),
	}
}

func ToResponses(users []models.User) []UserResponse {
	userResponses := make([]UserResponse, 0, len(users))

	for _, user := range users {
		userResponses = append(userResponses, *ToResponse(&user))
	}

	return userResponses
}

type PageResponse[T any] struct {
	Items []T `json:"items"`
	Total int `json:"total"`
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
}

func ToPageResponse(users []models.User, total int, page, limit int) *PageResponse[UserResponse] {
	userResponses := make([]UserResponse, 0, len(users))
	for _, user := range users {
		userResponses = append(userResponses, *ToResponse(&user))
	}
	return &PageResponse[UserResponse]{
		Items: userResponses,
		Total: total,
		Page:  page,
		Limit: limit,
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
