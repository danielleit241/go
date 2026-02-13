package main

import "fmt"

type Employee struct {
	Name string
	Age  int
	Role string
}

func (e *Employee) ToString() string {
	return fmt.Sprintf("Name: %s, Age: %d, Role: %s", e.Name, e.Age, e.Role)
}

type Subject struct {
	Code string
	Name string
}

func main() {
	drink := map[string]float32{
		"coffee": 2.50,
		"tea":    1.75,
		"juice":  3.00,
	}

	fmt.Printf("%+v\n", drink)

	number := make(map[string]int)

	fmt.Printf("Using make before appending values: %+v\n", number)

	number["one"] = 1
	number["two"] = 2
	number["three"] = 3

	fmt.Printf("After appending values: %+v\n", number)

	value, exists := number["four"]
	if exists {
		fmt.Printf("The value of 'four' is: %d\n", value)
	} else {
		fmt.Println("'four' does not exist in the map")
	}

	employees := map[string]Employee{
		"E001": {"Alice Johnson", 30, "Developer"},
		"E002": {"Bob Smith", 45, "Manager"},
		"E003": {"Charlie Brown", 28, "Designer"},
	}

	for key, value := range employees {
		fmt.Printf("Key: %s, Value: %+v\n", key, value.ToString())
	}

	studentSubjects := map[string][]Subject{
		"S001": {
			{Code: "MATH101", Name: "Calculus"},
			{Code: "PHY101", Name: "Physics"},
		},
		"S002": {
			{Code: "CHEM101", Name: "Chemistry"},
			{Code: "BIO101", Name: "Biology"},
		},
	}

	for studentID, subjects := range studentSubjects {
		fmt.Printf("Student ID: %s\n", studentID)
		for _, subject := range subjects {
			fmt.Printf("Subject Code: %s, Name: %s\n", subject.Code, subject.Name)
		}
	}
}
