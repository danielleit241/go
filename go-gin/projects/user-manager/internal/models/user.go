package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name" binding:"required,min=3,max=100"`
	Email    string    `json:"email" binding:"required,email"`
	Age      int       `json:"age" binding:"required,gte=0"`
	Password string    `json:"password" binding:"required,min=6"`
	Status   int       `json:"status" binding:"required,oneof=0 1"`
	Level    int       `json:"level" binding:"required,oneof=1 2"`
}
