package main

import "fmt"

func menu() {
	// Simple menu using for loop and switch-case
	for true {
		var choice int
		fmt.Println("Menu:")
		fmt.Println("1. Calculate Sum from 1 to N")
		fmt.Println("2. Calculate Fibonacci of N")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var n int
			fmt.Print("Enter a positive integer N: ")
			fmt.Scan(&n)
			result := sum1ToN(n)
			fmt.Printf("Sum from 1 to %d is %d\n", n, result)
		case 2:
			var n int
			fmt.Print("Enter a positive integer N: ")
			fmt.Scan(&n)
			result := fibonacciSlice(n)
			fmt.Printf("Fibonacci of %d is %d\n", n, result)
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func sum1ToN(n int) int {
	if n <= 0 {
		return 0
	}
	return n + sum1ToN(n-1)
}

func fibonacciSlice(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonacciSlice(n-1) + fibonacciSlice(n-2)
}
