package main

import (
	"fmt"
)

func menu() {
	for true {
		fmt.Println("Menu:")
		fmt.Println("1. Exercise 1: Print numbers from 0 to 100 excluding 6, 48, 75, and 89")
		fmt.Println("2. Exercise 2: Print multiples of 3 from 0 to 100, three per line")
		fmt.Println("3. Exercise 3: Print multiplication table for a given number")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice (1-4): ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			ex01()
		case 2:
			ex02()
		case 3:
			ex03()
		case 4:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func ex01() {
	counter := 0
	for i := 0; i <= 100; i++ {
		switch i {
		case 6, 48, 75, 89:
			continue
		default:
			fmt.Print(i, ", ")
		}
		counter++
	}
	fmt.Println("\nTotal numbers printed:", counter)
}

func ex02() {
	counter := 0

	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print(i)
			counter++

			if counter%3 == 0 {
				fmt.Println()
			} else {
				fmt.Print(", ")
			}
		}
	}

	fmt.Println("\nTotal numbers printed:", counter)
}

func ex03() {
	var number int
	fmt.Print("Enter a number to print its multiplication table: ")
	fmt.Scanln(&number)

	for i := 1; i <= 10; i++ {
		result := number * i
		fmt.Printf("%d x %d = %d\n", number, i, result)
	}
}
