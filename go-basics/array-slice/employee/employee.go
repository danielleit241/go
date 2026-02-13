package employee

import "fmt"

type Employee struct {
	ID       int
	FullName string
	Position string
	Salary   float64
}

func NewEmployee(id int, fullName, position string, salary float64) *Employee {
	return &Employee{
		ID:       id,
		FullName: fullName,
		Position: position,
		Salary:   salary,
	}
}

func (e *Employee) ToString() string {
	return fmt.Sprintf("ID: %d, Name: %s, Position: %s, Salary: %.2f",
		e.ID, e.FullName, e.Position, e.Salary)
}
