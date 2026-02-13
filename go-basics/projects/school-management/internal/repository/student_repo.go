package repository

import (
	"strings"

	"example.com/go/internal/entity"
)

type StudentReader interface {
	BaseReader[entity.Student]
	SearchByName(name string) ([]entity.Student, error)
}

type StudentWriter interface {
	BaseWriter[entity.Student]
}

type StudentRepository interface {
	StudentReader
	StudentWriter
}

type studentRepository struct {
	*baseRepository[entity.Student]
}

func NewStudentRepository() StudentRepository {
	return &studentRepository{
		baseRepository: newBaseRepository[entity.Student](),
	}
}

func (r *studentRepository) SearchByName(name string) ([]entity.Student, error) {
	var results []entity.Student
	all, err := r.GetAll()
	if err != nil {
		return nil, err
	}
	lowerName := strings.ToLower(name)
	for _, student := range all {
		fullName := student.FirstName + " " + student.LastName
		if strings.Contains(strings.ToLower(fullName), lowerName) {
			results = append(results, student)
		}
	}
	return results, nil
}
