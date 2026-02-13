package repository

import (
	"time"

	"example.com/go/entity"
)

type TransactionRepository interface {
	Repository[entity.Transaction]
	ReturnBook(id string) error
}

type inMemoryTransactionRepository struct {
	*inMemoryRepository[entity.Transaction]
}

func NewInMemoryTransactionRepository() *inMemoryTransactionRepository {
	return &inMemoryTransactionRepository{
		inMemoryRepository: NewInMemoryRepository[entity.Transaction](),
	}
}

func (r *inMemoryTransactionRepository) ReturnBook(id string) error {
	transaction, err := r.GetByID(id)
	if err != nil {
		return err
	}
	if transaction.ReturnDate != "" {
		return ErrBookAlreadyReturned
	}
	transaction.ReturnDate = time.Now().Format("2006-01-02")
	return r.Update(transaction)
}
