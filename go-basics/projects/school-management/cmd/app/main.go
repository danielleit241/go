package main

import (
	"github.com/danielleit241/internal/repository"
	"github.com/danielleit241/internal/service"
	"github.com/danielleit241/internal/ui"
)

func main() {
	studentRepo := repository.NewStudentRepository()
	lecturerRepo := repository.NewLecturerRepository()

	studentService := service.NewStudentService(studentRepo)
	lecturerService := service.NewLecturerService(lecturerRepo)

	menu := ui.NewMenu(studentService, lecturerService)
	menu.Run()
}
