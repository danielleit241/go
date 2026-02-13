package entity

type Lecturer struct {
	Person
	Department string
	Salary     float64
}

func (l Lecturer) GetID() string {
	return l.ID
}
