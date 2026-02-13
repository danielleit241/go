package entity

import "example.com/go/utils"

type User struct {
	ID    string
	Name  string
	Email string
}

func NewUser(name, email string) User {
	return User{
		ID:    utils.UUIDGenerator(),
		Name:  name,
		Email: email,
	}
}

func (u User) GetID() string {
	return u.ID
}
