package repository

import (
	"fmt"
)

var (
	ErrNotFound             = fmt.Errorf("entity not found")
	ErrAlreadyExists        = fmt.Errorf("entity already exists")
	ErrBookAlreadyAvailable = fmt.Errorf("book is already available")
	ErrBookNotAvailable     = fmt.Errorf("book is not available")
)

type BaseEntity interface {
	GetID() string
}

type Repository[T BaseEntity] interface {
	GetByID(id string) (T, error)
	Create(entity T) error
	Update(entity T) error
	Delete(id string) error
	GetAll() ([]T, error)
}

type inMemoryRepository[T BaseEntity] struct {
	items map[string]T
}

func NewInMemoryRepository[T BaseEntity]() *inMemoryRepository[T] {
	return &inMemoryRepository[T]{
		items: make(map[string]T),
	}
}

func (r *inMemoryRepository[T]) GetByID(id string) (T, error) {
	item, exists := r.items[id]
	if !exists {
		var zero T
		return zero, ErrNotFound
	}
	return item, nil
}

func (r *inMemoryRepository[T]) Create(entity T) error {
	id := entity.GetID()
	if _, exists := r.items[id]; exists {
		return ErrAlreadyExists
	}
	r.items[id] = entity
	return nil
}

func (r *inMemoryRepository[T]) Update(entity T) error {
	id := entity.GetID()
	if _, exists := r.items[id]; !exists {
		return ErrNotFound
	}
	r.items[id] = entity
	return nil
}

func (r *inMemoryRepository[T]) Delete(id string) error {
	if _, exists := r.items[id]; !exists {
		return ErrNotFound
	}
	delete(r.items, id)
	return nil
}

func (r *inMemoryRepository[T]) GetAll() ([]T, error) {
	var allItems []T
	for _, item := range r.items {
		allItems = append(allItems, item)
	}
	return allItems, nil
}
