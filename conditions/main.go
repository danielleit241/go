package main

import "fmt"

func main() {
	fmt.Println("Conditional Statements in Go")
	ifStatement()

	fmt.Println("\nSwitch Statement Example")
	switchStatement()
}

func switchStatement() {
	var day int
	fmt.Print("Enter a number (1-7) for the day of the week: ")
	fmt.Scanln(&day)
	switch day {
	case 1:
		println("Monday")
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	case 4:
		println("Thursday")
	case 5:
		println("Friday")
	case 6:
		println("Saturday")
	case 7:
		println("Sunday")
	default:
		println("Invalid day number!")
	}
}

func ifStatement() {
	var allowedAge int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&allowedAge)
	if allowedAge >= 18 {
		println("Access granted - you are old enough.")
	} else {
		println("Access denied - you are not old enough.")
	}

	var myGrade int
	fmt.Print("Enter your grade: ")
	fmt.Scanln(&myGrade)
	if myGrade >= 9 && myGrade <= 10 {
		println("You received an A.")
	} else if myGrade >= 8 && myGrade < 9 {
		println("You received a B.")
	} else if myGrade >= 6 && myGrade < 7 {
		println("You received a C.")
	} else if myGrade >= 5 && myGrade < 6 {
		println("You received a D.")
	} else {
		println("You received an F.")
	}
}
