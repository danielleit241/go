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

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100"`
	Age      int    `json:"age" binding:"required,gt=0"`
	Status   int    `json:"status" binding:"required,oneof=0 1"`
	Level    int    `json:"level" binding:"required,oneof=1 2"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" binding:"omitempty,min=2,max=100"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty,min=6,max=100"`
	Age      int    `json:"age" binding:"omitempty,gt=0"`
	Status   int    `json:"status" binding:"omitempty,oneof=0 1"`
	Level    int    `json:"level" binding:"omitempty,oneof=1 2"`
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

func (request *UserCreateRequest) ToEntity(req UserCreateRequest) models.User {
	return models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
		Status:   req.Status,
		Level:    req.Level,
	}
}

func (request *UserUpdateRequest) ToEntity(req UserUpdateRequest) models.User {
	return models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
		Status:   req.Status,
		Level:    req.Level,
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
