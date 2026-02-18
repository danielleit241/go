package ui

import (
	"fmt"
	"strings"

	"github.com/danielleit241/internal/entity"
	"github.com/danielleit241/internal/service"
)

type Menu struct {
	studentService  service.StudentService
	lecturerService service.LecturerService
}

func NewMenu(studentService service.StudentService, lecturerService service.LecturerService) *Menu {
	return &Menu{
		studentService:  studentService,
		lecturerService: lecturerService,
	}
}

func (m *Menu) Run() {
	for {
		fmt.Println("\n=== SCHOOL MANAGEMENT ===")
		fmt.Println("1. Student Management")
		fmt.Println("2. Lecturer Management")
		fmt.Println("0. Exit")

		choice := ReadInput("Select option: ")
		switch choice {
		case "1":
			m.runStudentMenu()
		case "2":
			m.runLecturerMenu()
		case "0":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (m *Menu) runStudentMenu() {
	for {
		fmt.Println("\n--- STUDENT MENU ---")
		fmt.Println("1. Create Student")
		fmt.Println("2. List Students")
		fmt.Println("3. Find Student By ID")
		fmt.Println("4. Search Student By Name")
		fmt.Println("5. Update Student")
		fmt.Println("6. Delete Student")
		fmt.Println("0. Back")

		choice := ReadInput("Select option: ")
		switch choice {
		case "1":
			m.createStudent()
		case "2":
			m.listStudents()
		case "3":
			m.findStudentByID()
		case "4":
			m.searchStudentByName()
		case "5":
			m.updateStudent()
		case "6":
			m.deleteStudent()
		case "0":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (m *Menu) runLecturerMenu() {
	for {
		fmt.Println("\n--- LECTURER MENU ---")
		fmt.Println("1. Create Lecturer")
		fmt.Println("2. List Lecturers")
		fmt.Println("3. Find Lecturer By ID")
		fmt.Println("4. Search Lecturer By Department")
		fmt.Println("5. Update Lecturer")
		fmt.Println("6. Delete Lecturer")
		fmt.Println("0. Back")

		choice := ReadInput("Select option: ")
		switch choice {
		case "1":
			m.createLecturer()
		case "2":
			m.listLecturers()
		case "3":
			m.findLecturerByID()
		case "4":
			m.searchLecturerByDepartment()
		case "5":
			m.updateLecturer()
		case "6":
			m.deleteLecturer()
		case "0":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (m *Menu) createStudent() {
	student, ok := inputStudent()
	if !ok {
		return
	}

	if err := m.studentService.Create(student); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Create student success")
}

func (m *Menu) listStudents() {
	students, err := m.studentService.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(students) == 0 {
		fmt.Println("No students found")
		return
	}

	for _, s := range students {
		fmt.Printf("ID=%s | Name=%s %s | Age=%d | Grade=%d | GPA=%.2f\n", s.ID, s.FirstName, s.LastName, s.Age, s.Grade, s.GPA)
	}
}

func (m *Menu) findStudentByID() {
	id := GetID("Student ID: ")
	student, err := m.studentService.GetByID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("ID=%s | Name=%s %s | Age=%d | Grade=%d | GPA=%.2f\n", student.ID, student.FirstName, student.LastName, student.Age, student.Grade, student.GPA)
}

func (m *Menu) searchStudentByName() {
	name := strings.TrimSpace(ReadInput("Name keyword: "))
	students, err := m.studentService.SearchByName(name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(students) == 0 {
		fmt.Println("No students found")
		return
	}
	for _, s := range students {
		fmt.Printf("ID=%s | Name=%s %s | Age=%d | Grade=%d | GPA=%.2f\n", s.ID, s.FirstName, s.LastName, s.Age, s.Grade, s.GPA)
	}
}

func (m *Menu) updateStudent() {
	student, ok := inputStudent()
	if !ok {
		return
	}

	if err := m.studentService.Update(student); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Update student success")
}

func (m *Menu) deleteStudent() {
	id := GetID("Student ID: ")
	if err := m.studentService.Delete(id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Delete student success")
}

func inputStudent() (entity.Student, bool) {
	id := GetID("ID: ")
	firstName := GetName("First name: ")
	lastName := GetName("Last name: ")
	age := GetAge("Age: ")
	grade, ok := readGrade("Grade: ")
	if !ok {
		return entity.Student{}, false
	}
	gpa, ok := readGPAFloat("GPA (X.XX): ")
	if !ok {
		return entity.Student{}, false
	}

	return entity.Student{
		Person: entity.Person{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Age:       age,
		},
		Grade: int16(grade),
		GPA:   float32(gpa),
	}, true
}

func readGrade(prompt string) (int, bool) {
	for {
		gradeStr := strings.TrimSpace(ReadInput(prompt))
		grade, err := ParseInt(gradeStr)
		if err != nil || grade < 1 || grade > 12 {
			fmt.Println("Error: Grade must be from 1 to 12")
			continue
		}
		return grade, true
	}
}

func readGPAFloat(prompt string) (float64, bool) {
	str := GetGPA(prompt)
	gpa, err := ParseFloat(str)
	if err != nil {
		fmt.Println("Error: Invalid GPA")
		return 0, false
	}
	return gpa, true
}

func (m *Menu) createLecturer() {
	lecturer := inputLecturer()
	if err := m.lecturerService.Create(lecturer); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Create lecturer success")
}

func (m *Menu) listLecturers() {
	lecturers, err := m.lecturerService.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(lecturers) == 0 {
		fmt.Println("No lecturers found")
		return
	}
	for _, l := range lecturers {
		fmt.Printf("ID=%s | Name=%s %s | Age=%d | Department=%s | Salary=%.2f\n", l.ID, l.FirstName, l.LastName, l.Age, l.Department, l.Salary)
	}
}

func (m *Menu) findLecturerByID() {
	id := GetID("Lecturer ID: ")
	lecturer, err := m.lecturerService.GetByID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("ID=%s | Name=%s %s | Age=%d | Department=%s | Salary=%.2f\n", lecturer.ID, lecturer.FirstName, lecturer.LastName, lecturer.Age, lecturer.Department, lecturer.Salary)
}

func (m *Menu) searchLecturerByDepartment() {
	department := GetDepartment("Department: ")
	lecturers, err := m.lecturerService.SearchByDepartment(department)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(lecturers) == 0 {
		fmt.Println("No lecturers found")
		return
	}
	for _, l := range lecturers {
		fmt.Printf("ID=%s | Name=%s %s | Age=%d | Department=%s | Salary=%.2f\n", l.ID, l.FirstName, l.LastName, l.Age, l.Department, l.Salary)
	}
}

func (m *Menu) updateLecturer() {
	lecturer := inputLecturer()
	if err := m.lecturerService.Update(lecturer); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Update lecturer success")
}

func (m *Menu) deleteLecturer() {
	id := GetID("Lecturer ID: ")
	if err := m.lecturerService.Delete(id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Delete lecturer success")
}

func inputLecturer() entity.Lecturer {
	id := GetID("ID: ")
	firstName := GetName("First name: ")
	lastName := GetName("Last name: ")
	age := GetAge("Age: ")
	department := GetDepartment("Department: ")
	salary := GetSalary("Salary: ")

	return entity.Lecturer{
		Person: entity.Person{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Age:       age,
		},
		Department: department,
		Salary:     salary,
	}
}
