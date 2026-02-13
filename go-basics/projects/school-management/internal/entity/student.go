package entity

type Student struct {
	Person
	Grade int16
	GPA   float32
}

func (s Student) GetID() string {
	return s.ID
}
