package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email" `
	Age      int       `json:"age"`
	Password string    `json:"password"`
	Status   int       `json:"status"`
	Level    int       `json:"level"`
}
