package ui

import (
	"fmt"

	"example.com/go/service"
	"example.com/go/utils"
)

type Menu struct {
	LibraryService service.LibraryService
}

func NewMenu(libraryService service.LibraryService) *Menu {
	return &Menu{
		LibraryService: libraryService,
	}
}

func (m *Menu) Show() {
	for {
		m.printMenu()
		choice := utils.ReadInput("Choose an option: ")

		switch choice {
		case "1":
			m.addBook()
		case "2":
			m.registerUser()
		case "3":
			m.borrowBook()
		case "4":
			m.returnBook()
		case "5":
			m.listBooks()
		case "6":
			m.searchBooksByTitle()
		case "7":
			m.listUsers()
		case "8":
			m.listTransactions()
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option")
		}

		fmt.Println()
	}
}

func (m *Menu) printMenu() {
	fmt.Println("===== Library Management =====")
	fmt.Println("1. Add book")
	fmt.Println("2. Register user")
	fmt.Println("3. Borrow book")
	fmt.Println("4. Return book")
	fmt.Println("5. List books")
	fmt.Println("6. Search books by title")
	fmt.Println("7. List users")
	fmt.Println("8. List transactions")
	fmt.Println("0. Exit")
}

func (m *Menu) addBook() {
	title := utils.GetBookTitle("Book title: ")
	author := utils.GetBookAuthor("Book author: ")

	if err := m.LibraryService.AddBook(title, author); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Book added successfully")
}

func (m *Menu) registerUser() {
	name := utils.GetUserName("User name: ")
	email := utils.GetUserEmail("User email: ")

	if err := m.LibraryService.RegisterUser(name, email); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User registered successfully")
}

func (m *Menu) borrowBook() {
	bookID := utils.GetBookID("Book ID: ")
	userID := utils.GetUserID("User ID: ")

	if err := m.LibraryService.BorrowBook(bookID, userID); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Book borrowed successfully")
}

func (m *Menu) returnBook() {
	bookID := utils.GetBookID("Book ID: ")
	userID := utils.GetUserID("User ID: ")

	if err := m.LibraryService.ReturnBook(bookID, userID); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Book returned successfully")
}

func (m *Menu) listBooks() {
	books, err := m.LibraryService.GetBooks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books found")
		return
	}

	for _, book := range books {
		fmt.Printf("ID: %s | Title: %s | Author: %s | Available: %t\n", book.ID, book.Title, book.Author, book.IsAvailable)
	}
}

func (m *Menu) searchBooksByTitle() {
	title := utils.GetBookTitle("Title to search: ")
	books, err := m.LibraryService.SearchBooksByTitle(title)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books found")
		return
	}

	for _, book := range books {
		fmt.Printf("ID: %s | Title: %s | Author: %s | Available: %t\n", book.ID, book.Title, book.Author, book.IsAvailable)
	}
}

func (m *Menu) listUsers() {
	users, err := m.LibraryService.GetUsers()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(users) == 0 {
		fmt.Println("No users found")
		return
	}

	for _, user := range users {
		fmt.Printf("ID: %s | Name: %s | Email: %s\n", user.ID, user.Name, user.Email)
	}
}

func (m *Menu) listTransactions() {
	transactions, err := m.LibraryService.GetTransactions()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(transactions) == 0 {
		fmt.Println("No transactions found")
		return
	}

	for _, transaction := range transactions {
		fmt.Printf("ID: %s | BookID: %s | UserID: %s | BorrowDate: %s | ReturnDate: %s\n", transaction.ID, transaction.BookID, transaction.UserID, transaction.BorrowDate, transaction.ReturnDate)
	}
}
