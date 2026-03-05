package v1dto

import (
	"github.com/danielleit241/internal/db/sqlc"
	"github.com/danielleit241/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Age          *int      `json:"age"`
	Status       string    `json:"status"`
	Role         string    `json:"role"`
	CreatedAtUtc string    `json:"created_at_utc"`
	UpdatedAtUtc string    `json:"updated_at_utc"`
	DeletedAtUtc *string   `json:"deleted_at_utc"`
}

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100"`
	Age      *int   `json:"age" binding:"omitempty,gt=0"`
	Status   int    `json:"status" binding:"required,oneof=0 1"`
	Role     int    `json:"role" binding:"required,oneof=0 1 2"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" binding:"omitempty,min=2,max=100"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty,min=6,max=100"`
	Age      *int   `json:"age" binding:"omitempty,gt=0"`
	Status   int    `json:"status" binding:"omitempty,oneof=0 1"`
	Role     int    `json:"role" binding:"omitempty,oneof=0 1 2"`
}

func ToResponse(user sqlc.User) *UserResponse {
	return &UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Age:          utils.ConvertInt32PtrToInt(user.Age),
		Status:       GetStringStatus(int(user.Status)),
		Role:         GetStringLevel(int(user.Role)),
		CreatedAtUtc: user.CreatedAtUtc.Format(utils.DateTimeFormat),
		UpdatedAtUtc: user.UpdatedAtUtc.Format(utils.DateTimeFormat),
		DeletedAtUtc: formatTimestamptz(user.DeletedAtUtc),
	}
}

func formatTimestamptz(t pgtype.Timestamptz) *string {
	if !t.Valid {
		return nil
	}
	formatted := t.Time.Format(utils.DateTimeFormat)
	return &formatted
}

func ToResponses(users []sqlc.User) []UserResponse {
	responses := make([]UserResponse, len(users))
	for i, user := range users {
		responses[i] = *ToResponse(user)
	}
	return responses
}

func (request *UserCreateRequest) ToCreateEntity() sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Age:      utils.ConvertToInt32Ptr(request.Age),
		Status:   int32(request.Status),
		Role:     int32(request.Role),
	}
}

func (request *UserUpdateRequest) ToUpdateEntity() {
}

func GetStringStatus(status int) string {
	if status == 1 {
		return "active"
	}
	return "inactive"
}

func GetStringLevel(level int) string {
	switch level {
	case 1:
		return "moderator"
	case 2:
		return "admin"
	default:
		return "user"
	}
}
