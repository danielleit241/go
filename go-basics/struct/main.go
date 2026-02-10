package main

import (
	"fmt"

	"example.com/go/model"
)

func main() {
	student := model.Student{StudentId: 1, Name: "Alice", Age: 20, IsEnrolled: false}
	student.DisplayInfo()
	fmt.Println()
	fmt.Println("Clearing student data...")
	student.Clear()
	student.DisplayInfo()

	//student.checkEnroll() // This will cause a compile-time error since checkEnroll is unexported
}
