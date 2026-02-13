package service

import (
	"time"

	"example.com/go/entity"
	"example.com/go/repository"
)

type LibraryReader interface {
	GetBooks() ([]entity.Book, error)
	SearchBooksByTitle(title string) ([]entity.Book, error)
	GetTransactions() ([]entity.Transaction, error)
	GetUsers() ([]entity.User, error)
}

type LibraryWriter interface {
	AddBook(title, author string) error
	RegisterUser(name, email string) error
	BorrowBook(bookID, userID string) error
	ReturnBook(bookID, userID string) error
}

type LibraryService interface {
	LibraryReader
	LibraryWriter
}

type libraryService struct {
	BookRepo        repository.BookRepository
	UserRepo        repository.UserRepository
	TransactionRepo repository.TransactionRepository
}

func NewLibraryService(bookRepo repository.BookRepository, userRepo repository.UserRepository, transactionRepo repository.TransactionRepository) *libraryService {
	return &libraryService{
		BookRepo:        bookRepo,
		UserRepo:        userRepo,
		TransactionRepo: transactionRepo,
	}
}

func (s *libraryService) GetBooks() ([]entity.Book, error) {
	return s.BookRepo.GetAll()
}

func (s *libraryService) SearchBooksByTitle(title string) ([]entity.Book, error) {
	return s.BookRepo.FindBookByTitle(title)
}

func (s *libraryService) GetTransactions() ([]entity.Transaction, error) {
	return s.TransactionRepo.GetAll()
}

func (s *libraryService) GetUsers() ([]entity.User, error) {
	return s.UserRepo.GetAll()
}

func (s *libraryService) AddBook(title, author string) error {
	book := entity.NewBook(title, author)
	return s.BookRepo.Create(book)
}

func (s *libraryService) RegisterUser(name, email string) error {
	user := entity.NewUser(name, email)
	return s.UserRepo.Create(user)
}

func (s *libraryService) BorrowBook(bookID, userID string) error {
	err := s.BookRepo.BorrowBook(bookID)
	if err != nil {
		return err
	}

	borrowDate := time.Now().Format("2006-01-02")
	transaction := entity.NewTransaction(bookID, userID, borrowDate, "")
	return s.TransactionRepo.Create(transaction)
}
func (s *libraryService) ReturnBook(bookID, userID string) error {
	err := s.BookRepo.ReturnBook(bookID)
	if err != nil {
		return err
	}
	returnDate := time.Now().Format("2006-01-02")
	transaction := entity.NewTransaction(bookID, userID, "", returnDate)
	return s.TransactionRepo.Create(transaction)
}
