package main

import (
	"example.com/go/internal/repository"
	"example.com/go/internal/service"
	"example.com/go/internal/ui"
)

func main() {
	studentRepo := repository.NewStudentRepository()
	lecturerRepo := repository.NewLecturerRepository()

	studentService := service.NewStudentService(studentRepo)
	lecturerService := service.NewLecturerService(lecturerRepo)

	menu := ui.NewMenu(studentService, lecturerService)
	menu.Run()
}
