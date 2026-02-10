package model

import "fmt"

type Student struct {
	StudentId  int
	Name       string
	Age        int
	IsEnrolled bool
}

// receiver function = methods
// public methods start with uppercase letter
func (s Student) DisplayInfo() {
	fmt.Printf("Student ID: %d\n", s.StudentId)
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Printf("Enrolled: %t\n", s.IsEnrolled)
}

func (s *Student) Enroll() {
	if !s.checkEnroll() {
		fmt.Println("Student is now enrolled.")
	}
	s.IsEnrolled = true
}

func (s *Student) Clear() {
	s.StudentId = 0
	s.Name = ""
	s.Age = 0
	s.IsEnrolled = false
}

// private methods start with lowercase letter
func (s *Student) checkEnroll() bool {
	return s.IsEnrolled
}
