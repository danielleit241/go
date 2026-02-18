package main

import (
	"github.com/danielleit241/repository"
	"github.com/danielleit241/service"
	"github.com/danielleit241/ui"
)

func main() {
	bookRepo := repository.NewInMemoryBookRepository()
	userRepo := repository.NewInMemoryUserRepository()
	transactionRepo := repository.NewInMemoryTransactionRepository()

	libraryService := service.NewLibraryService(bookRepo, userRepo, transactionRepo)
	menu := ui.NewMenu(libraryService)
	menu.Show()
}
