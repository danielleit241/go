package entity

import "github.com/danielleit241/utils"

type Book struct {
	ID          string
	Title       string
	Author      string
	IsAvailable bool
}

func NewBook(title, author string) Book {
	return Book{
		ID:          utils.UUIDGenerator(),
		Title:       title,
		Author:      author,
		IsAvailable: true,
	}
}

func (b Book) GetID() string {
	return b.ID
}
