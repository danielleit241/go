package main

import "fmt"

func main() {
	// fmt.Println("Conditional Statements in Go")
	// ifStatement()

	fmt.Println("\nSwitch Statement Example")
	switchStatement()
}

func switchStatement() {
	var day int
	fmt.Print("Enter a number (1-7) for the day of the week: ")
	fmt.Scanln(&day)
	switch day {
	case 1:
		fmt.Println("Monday")
		// No break needed; Go automatically breaks after each case
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6, 7: // Multiple cases can be handled together
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day number!")
	}
}

func ifStatement() {
	var allowedAge int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&allowedAge)
	if allowedAge >= 18 {
		fmt.Println("Access granted - you are old enough.")
	} else {
		fmt.Println("Access denied - you are not old enough.")
	}

	var myGrade int
	fmt.Print("Enter your grade: ")
	fmt.Scanln(&myGrade)
	if myGrade >= 9 && myGrade <= 10 {
		fmt.Println("You received an A.")
	} else if myGrade >= 8 && myGrade < 9 {
		fmt.Println("You received a B.")
	} else if myGrade >= 6 && myGrade < 7 {
		fmt.Println("You received a C.")
	} else if myGrade >= 5 && myGrade < 6 {
		fmt.Println("You received a D.")
	} else {
		fmt.Println("You received an F.")
	}
}
