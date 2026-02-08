package main

import "fmt"

func main() {
	// printExamples()
	// scanExamples()
	sprintExamples()
}

func sprintExamples() {
	name := "Alice"
	age := 30
	message1 := fmt.Sprint("Name: ", name, ", Age: ", age)
	fmt.Println("Using Sprint:", message1)

	message2 := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println("Using Sprintf:", message2)

	message3 := fmt.Sprintln("Name:", name, ", Age:", age)
	fmt.Print("Using Sprintln:", message3)
}

func scanExamples() {
	// Scan:
	// var firstName, lastName string
	// fmt.Print("Please enter your full name: ")
	// fmt.Scan(&firstName, &lastName)
	// fmt.Printf("Hello, %s %s!\n", firstName, lastName)

	//Scanln:
	// var firstName, lastName string
	// fmt.Print("Please enter your full name: ")
	// fmt.Scanln(&firstName, &lastName) // Using enter key to end input
	// fmt.Printf("Hello, %s %s!\n", firstName, lastName)

	// var firstName, lastName string
	// fmt.Print("Please enter your first name: ")
	// fmt.Scanln(&firstName)
	// fmt.Print("Please enter your last name: ")
	// fmt.Scanln(&lastName)
	// fmt.Printf("Hello, %s %s!\n", firstName, lastName)

	//Scanf:
	// var name string
	// var age int
	// fmt.Print("Please enter your name: ")
	// fmt.Scanf("%s", &name)
	// fmt.Scanln() // Catch buffer newline
	// fmt.Print("Please enter your age: ")
	// fmt.Scanf("%d", &age)
	// fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}

func printExamples() {
	// fmt.Print("Print without a newline")
	// fmt.Println()                       // Newline
	// fmt.Println("Print with a newline") // Print with a newline

	// const PI = 3.14
	// var name = "Circle"
	// fmt.Printf("Value of PI: %.2f in %s\n", PI, name) // Formatted print
}
