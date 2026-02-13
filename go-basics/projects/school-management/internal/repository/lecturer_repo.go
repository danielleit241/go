package repository

import "example.com/go/internal/entity"

type LecturerReader interface {
	BaseReader[entity.Lecturer]
	SearchByDepartment(department string) ([]entity.Lecturer, error)
}

type LecturerWriter interface {
	BaseWriter[entity.Lecturer]
}

type LecturerRepository interface {
	LecturerReader
	LecturerWriter
}

type lecturerRepository struct {
	*baseRepository[entity.Lecturer]
}

func NewLecturerRepository() LecturerRepository {
	return &lecturerRepository{
		baseRepository: newBaseRepository[entity.Lecturer](),
	}
}

func (r *lecturerRepository) SearchByDepartment(department string) ([]entity.Lecturer, error) {
	var results []entity.Lecturer
	all, err := r.GetAll()
	if err != nil {
		return nil, err
	}
	for _, lecturer := range all {
		if lecturer.Department == department {
			results = append(results, lecturer)
		}
	}
	return results, nil
}
