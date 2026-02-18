package service

import (
	"strings"

	"github.com/danielleit241/internal/entity"
	"github.com/danielleit241/internal/repository"
)

type LecturerService interface {
	Create(lecturer entity.Lecturer) error
	GetByID(id string) (entity.Lecturer, error)
	Update(lecturer entity.Lecturer) error
	Delete(id string) error
	GetAll() ([]entity.Lecturer, error)
	SearchByDepartment(department string) ([]entity.Lecturer, error)
}

type lecturerService struct {
	reader repository.LecturerReader
	writer repository.LecturerWriter
}

func NewLecturerService(repo repository.LecturerRepository) LecturerService {
	return &lecturerService{
		reader: repo,
		writer: repo,
	}
}

func NewLecturerServiceWithPorts(reader repository.LecturerReader, writer repository.LecturerWriter) LecturerService {
	return &lecturerService{
		reader: reader,
		writer: writer,
	}
}

func (s *lecturerService) Create(lecturer entity.Lecturer) error {
	if err := validateLecturer(lecturer); err != nil {
		return err
	}
	return s.writer.Create(lecturer)
}

func (s *lecturerService) GetByID(id string) (entity.Lecturer, error) {
	if err := validateID(id); err != nil {
		return entity.Lecturer{}, err
	}
	return s.reader.GetByID(id)
}

func (s *lecturerService) Update(lecturer entity.Lecturer) error {
	if err := validateLecturer(lecturer); err != nil {
		return err
	}
	return s.writer.Update(lecturer)
}

func (s *lecturerService) Delete(id string) error {
	if err := validateID(id); err != nil {
		return err
	}
	return s.writer.Delete(id)
}

func (s *lecturerService) GetAll() ([]entity.Lecturer, error) {
	return s.reader.GetAll()
}

func (s *lecturerService) SearchByDepartment(department string) ([]entity.Lecturer, error) {
	if strings.TrimSpace(department) == "" {
		return nil, repository.ErrInvalidInput
	}
	return s.reader.SearchByDepartment(department)
}

func validateLecturer(lecturer entity.Lecturer) error {
	if err := validateID(lecturer.ID); err != nil {
		return err
	}
	if strings.TrimSpace(lecturer.FirstName) == "" || strings.TrimSpace(lecturer.LastName) == "" {
		return repository.ErrInvalidInput
	}
	if strings.TrimSpace(lecturer.Department) == "" {
		return repository.ErrInvalidInput
	}
	if lecturer.Age <= 0 {
		return repository.ErrInvalidInput
	}
	if lecturer.Salary < 0 {
		return repository.ErrInvalidInput
	}
	return nil
}
