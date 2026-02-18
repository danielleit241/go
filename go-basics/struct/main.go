package main

import (
	"encoding/json"
	"fmt"

	"github.com/danielleit241/model"
)

func main() {
	// studentStructExample()
	ex01()
}

func ex01() {
	rect := model.NewRectangle(5, 10)
	if rect == nil {
		fmt.Println("Invalid rectangle dimensions.")
		return
	}
	rect.PrintDetails()
	fmt.Printf("Rectangle Area: %.2f\n", rect.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", rect.Perimeter())
}

func studentStructExample() {
	student := model.NewStudent(1, "John Doe", 21)
	if student == nil {
		fmt.Println("Invalid student data.")
		return
	}
	student.DisplayInfo()
	student.Enroll()
	student.DisplayInfo()

	var jsonData, err = json.Marshal(student)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	} else {
		fmt.Println("JSON data:", string(jsonData))
	}

	//student.checkEnroll() // This will cause a compile-time error since checkEnroll is unexported
}
