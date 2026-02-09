package main

import "fmt"

func main() {
	// printExample()
	// scanExample()
	// sprintExample()
	formattingVerbsExample()
}

func formattingVerbsExample() {
	name := "Alice"
	age := 30
	height := 5.7
	isGraduated := true
	percent := 85.5

	fmt.Println("Data Type:")
	fmt.Printf("Data type of name is: %T\n", name) // %T for type
	fmt.Printf("Data type of age is: %T\n", age)
	fmt.Printf("Data type of height is: %T\n", height)
	fmt.Printf("Data type of isGraduated is: %T\n", isGraduated)
	fmt.Printf("Data type of percent is: %T\n", percent)

	fmt.Println("---")
	fmt.Println("Value:")
	fmt.Printf("Value of name is: %v\n", name)                      // %v for value
	fmt.Printf("My name is %s and I am %d years old.\n", name, age) // %s for string, %d for integer
	fmt.Printf("Height: %.2f feet\n", height)                       // %.2f for float with 2 decimal places
	fmt.Printf("Graduated: %t\n", isGraduated)                      // %t for boolean
	fmt.Printf("Score: %.1f%%\n", percent)                          // %% for literal percent sign
}

func sprintExample() {
	name := "Alice"
	age := 30
	message1 := fmt.Sprint("Name: ", name, ", Age: ", age)
	fmt.Println("Using Sprint:", message1)

	message2 := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println("Using Sprintf:", message2)

	message3 := fmt.Sprintln("Name:", name, ", Age:", age)
	fmt.Print("Using Sprintln:", message3)
}

func scanExample() {
	// Scan:
	var firstName1, lastName1 string
	fmt.Print("Please enter your full name: ")
	fmt.Scan(&firstName1, &lastName1)
	fmt.Printf("Hello, %s %s!\n", firstName1, lastName1)

	//Scanln:
	// var firstName2, lastName2 string
	// fmt.Print("Please enter your full name: ")
	// fmt.Scanln(&firstName2, &lastName2) // Using enter key to end input
	// fmt.Printf("Hello, %s %s!\n", firstName2, lastName2)

	var firstName2, lastName2 string
	fmt.Print("Please enter your first name: ")
	fmt.Scanln(&firstName2)
	fmt.Print("Please enter your last name: ")
	fmt.Scanln(&lastName2)
	fmt.Printf("Hello, %s %s!\n", firstName2, lastName2)

	//Scanf:
	var name string
	var age int
	fmt.Print("Please enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Scanln() // Catch buffer newline
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%d", &age)
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}

func printExample() {
	fmt.Print("Print without a newline")
	fmt.Println()                       // Newline
	fmt.Println("Print with a newline") // Print with a newline

	const PI = 3.14
	var name = "Circle"
	fmt.Printf("Value of PI: %.2f in %s\n", PI, name) // Formatted print
}
