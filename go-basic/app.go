package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var address string = "123 Main St"

var (
	courseId   int
	courseName string = "Go Programming"
)

// course := "Go Programming" // shorthand declaration at package level is not allowed

func main() {
	// fmt.Println("Hello, Go!")
	// randomUser()

	// fmt.Println("Address variable example:", address)
	// variableExamples()

	// courseId = 101
	// courseName = "Advanced Go Programming"
	// fmt.Println("Course ID:", courseId)
	// fmt.Println("Course Name:", courseName)

	var fullName string = "John Doe"
	fmt.Println("Full Name variable example:", fullName)

	fullName = "Jane Smith"
	fmt.Println("Updated Full Name variable example:", fullName)

	fmt.Print("Enter your full name: ")
	// fmt.Scan(&fullName) // Note: This will only capture input until the first space
	// fmt.Println("Scanned Full Name variable example:", fullName)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fullName = scanner.Text()
	}
	fmt.Println("Scanned Full Name variable example:", fullName)

	var age int
	fmt.Print("Enter your age: ")
	if scanner.Scan() {
		ageInput := scanner.Text()
		var err error
		age, err = strconv.Atoi(ageInput)
		if err != nil {
			fmt.Println("Vui lòng nhập số hợp lệ!")
		}
	}
	fmt.Println("Scanned Age variable example:", age)
}

func variableExamples() {
	var number int8 = 10 // 1 byte
	fmt.Println("Variable example:", number)

	var number2 int16 = 300 // 2 bytes
	fmt.Println("Another variable example:", number2)

	var number3 int32 = 70000 // 4 bytes
	fmt.Println("Yet another variable example:", number3)

	var number4 int64 = 5000000000 // 8 bytes
	fmt.Println("And another variable example:", number4)

	var floatNum float32 = 5.75 // 4 bytes
	fmt.Println("Float variable example:", floatNum)

	var floatNum2 float64 = 19.99 // 8 bytes
	fmt.Println("Another float variable example:", floatNum2)

	var age int
	age = 10
	fmt.Println("Age variable example:", age)

	var year = 2024 // type inferred as int
	fmt.Println("Year variable example:", year)

	// Only inside functions
	score := 95 // shorthand declaration, type inferred as int
	fmt.Println("Score variable example:", score)

	// Multiple variable declaration
	var math1, science1, english1 int = 90, 85, 88
	fmt.Println("Multiple variable example:", math1, science1, english1)

	math2, science2, english2 := 92, 89, 84
	fmt.Println("Another multiple variable example:", math2, science2, english2)
}
