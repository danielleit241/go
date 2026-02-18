package service

import (
	"strings"

	"github.com/danielleit241/internal/entity"
	"github.com/danielleit241/internal/repository"
)

type StudentService interface {
	Create(student entity.Student) error
	GetByID(id string) (entity.Student, error)
	Update(student entity.Student) error
	Delete(id string) error
	GetAll() ([]entity.Student, error)
	SearchByName(name string) ([]entity.Student, error)
}

type studentService struct {
	reader repository.StudentReader
	writer repository.StudentWriter
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{
		reader: repo,
		writer: repo,
	}
}

func NewStudentServiceWithPorts(reader repository.StudentReader, writer repository.StudentWriter) StudentService {
	return &studentService{
		reader: reader,
		writer: writer,
	}
}

func (s *studentService) Create(student entity.Student) error {
	if err := validateStudent(student); err != nil {
		return err
	}
	return s.writer.Create(student)
}

func (s *studentService) GetByID(id string) (entity.Student, error) {
	if err := validateID(id); err != nil {
		return entity.Student{}, err
	}
	return s.reader.GetByID(id)
}

func (s *studentService) Update(student entity.Student) error {
	if err := validateStudent(student); err != nil {
		return err
	}
	return s.writer.Update(student)
}

func (s *studentService) Delete(id string) error {
	if err := validateID(id); err != nil {
		return err
	}
	return s.writer.Delete(id)
}

func (s *studentService) GetAll() ([]entity.Student, error) {
	return s.reader.GetAll()
}

func (s *studentService) SearchByName(name string) ([]entity.Student, error) {
	if strings.TrimSpace(name) == "" {
		return nil, repository.ErrInvalidInput
	}
	return s.reader.SearchByName(name)
}

func validateID(id string) error {
	if strings.TrimSpace(id) == "" {
		return repository.ErrInvalidInput
	}
	return nil
}

func validateStudent(student entity.Student) error {
	if err := validateID(student.ID); err != nil {
		return err
	}
	if strings.TrimSpace(student.FirstName) == "" || strings.TrimSpace(student.LastName) == "" {
		return repository.ErrInvalidInput
	}
	if student.Age <= 0 {
		return repository.ErrInvalidInput
	}
	return nil
}
