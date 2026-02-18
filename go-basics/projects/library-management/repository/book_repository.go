package repository

import "github.com/danielleit241/entity"

type BookRepository interface {
	Repository[entity.Book]
	FindBookByTitle(title string) ([]entity.Book, error)
	BorrowBook(id string) error
}

type inMemoryBookRepository struct {
	*inMemoryRepository[entity.Book]
}

func NewInMemoryBookRepository() *inMemoryBookRepository {
	return &inMemoryBookRepository{
		inMemoryRepository: NewInMemoryRepository[entity.Book](),
	}
}

func (r *inMemoryBookRepository) FindBookByTitle(title string) ([]entity.Book, error) {
	var results []entity.Book
	for _, book := range r.items {
		if book.Title == title {
			results = append(results, book)
		}
	}
	return results, nil
}

func (r *inMemoryBookRepository) BorrowBook(id string) error {
	book, err := r.GetByID(id)
	if err != nil {
		return err
	}
	if !book.IsAvailable {
		return ErrBookNotAvailable
	}
	book.IsAvailable = false
	return r.Update(book)
}
