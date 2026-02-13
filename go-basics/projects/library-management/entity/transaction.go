package entity

import "example.com/go/utils"

type Transaction struct {
	ID         string
	BookID     string
	UserID     string
	BorrowDate string
	ReturnDate string
}

func NewTransaction(bookID, userID, borrowDate, returnDate string) Transaction {
	return Transaction{
		ID:         utils.UUIDGenerator(),
		BookID:     bookID,
		UserID:     userID,
		BorrowDate: borrowDate,
		ReturnDate: returnDate,
	}
}

func (t Transaction) GetID() string {
	return t.ID
}
