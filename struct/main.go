package main

import "fmt"

type Student struct {
	studentId  int
	name       string
	age        int
	isEnrolled bool
}

func main() {
	// alice := Student{
	// 	studentId:  1,
	// 	name:       "Alice",
	// 	age:        20,
	// 	isEnrolled: true,
	// }

	// johnDoe := Student{
	// 	studentId:  2,
	// 	name:       "John Doe",
	// 	age:        22,
	// 	isEnrolled: false,
	// }

	// printStudentDetails(alice)
	// printStudentDetails(johnDoe)

	student1 := createStudent(101, "Bob", 21, true)
	printStudentDetails(student1)
	fmt.Println("Updating student name...")
	updateStudentName(&student1, "Robert")
	printStudentDetails(student1)
}

func updateStudentName(s *Student, newName string) {
	s.name = newName
}

func createStudent(id int, name string, age int, enrolled bool) Student {
	return Student{
		studentId:  id,
		name:       name,
		age:        age,
		isEnrolled: enrolled,
	}
}

func printStudentDetails(s Student) {
	fmt.Printf("ID: %d\n", s.studentId)
	fmt.Printf("Name: %s\n", s.name)
	fmt.Printf("Age: %d\n", s.age)
	fmt.Printf("Enrolled: %t\n", s.isEnrolled)
}
