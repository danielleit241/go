package repository

import (
	"errors"
	"slices"
)

var (
	ErrNotFound     = errors.New("Bản ghi không tồn tại")
	ErrDuplicateID  = errors.New("Id đã tồn tại")
	ErrInvalidInput = errors.New("Dữ liệu đầu vào không hợp lệ")
)

type HasID interface {
	GetID() string
}

type BaseReader[T HasID] interface {
	GetByID(id string) (T, error)
	GetAll() ([]T, error)
}

type BaseWriter[T HasID] interface {
	Create(item T) error
	Update(item T) error
	Delete(id string) error
}

type BaseRepository[T HasID] interface {
	BaseReader[T]
	BaseWriter[T]
}

type baseRepository[T HasID] struct {
	items []T
}

func newBaseRepository[T HasID]() *baseRepository[T] {
	return &baseRepository[T]{
		items: make([]T, 0),
	}
}

func (r *baseRepository[T]) findIndex(id string) int {
	for i := range r.items {
		if r.items[i].GetID() == id {
			return i
		}
	}
	return -1
}

func (r *baseRepository[T]) GetAll() ([]T, error) {
	return slices.Clone(r.items), nil
}

func (r *baseRepository[T]) GetByID(id string) (T, error) {
	index := r.findIndex(id)
	if index == -1 {
		var zero T
		return zero, ErrNotFound
	}
	return r.items[index], nil
}

func (r *baseRepository[T]) Create(item T) error {
	if r.findIndex(item.GetID()) != -1 {
		return ErrDuplicateID
	}
	r.items = append(r.items, item)
	return nil
}

func (r *baseRepository[T]) Update(item T) error {
	index := r.findIndex(item.GetID())
	if index == -1 {
		return ErrNotFound
	}
	r.items[index] = item
	return nil
}

func (r *baseRepository[T]) Delete(id string) error {
	index := r.findIndex(id)
	if index == -1 {
		return ErrNotFound
	}
	r.items = slices.Delete(r.items, index, index+1)
	return nil
}
