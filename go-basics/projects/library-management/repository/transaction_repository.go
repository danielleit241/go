package repository

import "example.com/go/entity"

type TransactionRepository interface {
	Repository[entity.Transaction]
}

type inMemoryTransactionRepository struct {
	*inMemoryRepository[entity.Transaction]
}

func NewInMemoryTransactionRepository() *inMemoryTransactionRepository {
	return &inMemoryTransactionRepository{
		inMemoryRepository: NewInMemoryRepository[entity.Transaction](),
	}
}
