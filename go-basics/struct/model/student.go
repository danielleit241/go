package model

import "fmt"

type Student struct {
	StudentId int `json:"student_identifier_number" desc:"The unique identifier for a student"` //public field with JSON tag
	// struct tags provide metadata about the fields
	Name       string `json:"name"`
	Age        int    `json:"age"`
	isEnrolled bool   // private field
}

func NewStudent(studentId int, name string, age int) *Student {
	if studentId <= 0 || name == "" || age < 0 {
		return nil
	}
	return &Student{
		StudentId: studentId,
		Name:      name,
		Age:       age,
	}
}

// receiver function = methods
// public methods start with uppercase letter
func (s Student) DisplayInfo() {
	fmt.Printf("Student ID: %d\n", s.StudentId)
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Printf("Enrolled: %t\n", s.isEnrolled)
}

func (s *Student) Enroll() {
	if !s.checkEnroll() {
		fmt.Println("Student is now enrolled.")
	}
	s.isEnrolled = true
}

func (s *Student) Clear() {
	s.StudentId = 0
	s.Name = ""
	s.Age = 0
	s.isEnrolled = false
}

// private methods start with lowercase letter
func (s *Student) checkEnroll() bool {
	return s.isEnrolled
}
