package main

import (
	"example.com/go/repository"
	"example.com/go/service"
	"example.com/go/ui"
)

func main() {
	bookRepo := repository.NewInMemoryBookRepository()
	userRepo := repository.NewInMemoryUserRepository()
	transactionRepo := repository.NewInMemoryTransactionRepository()

	libraryService := service.NewLibraryService(bookRepo, userRepo, transactionRepo)
	menu := ui.NewMenu(libraryService)
	menu.Show()
}
