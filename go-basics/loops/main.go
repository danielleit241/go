package main

import "fmt"

func main() {
	// fmt.Println("For Loop Example")
	// forILoop()

	fmt.Println("\nMenu Example")
	menu()
}

func forILoop() {
	//Print first 5 numbers in the loop
	fmt.Println("First 5 numbers:")
	for i := 1; i <= 50; i++ {
		if i > 5 {
			break
		}
		fmt.Println("Number:", i)
	}

	fmt.Println("Using continue statement:")
	for j := 1; j <= 5; j++ {
		if j == 3 {
			continue
		}
		fmt.Println("Number with continue:", j)
	}

	fmt.Println("Print odd numbers from 1 to 10")
	//Print even numbers from 2 to 10
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even Number:", i)
		}
	}

	for i := 0; i < 10; i += 2 {
		fmt.Println("Even Number (step 2):", i)
	}

	fmt.Println("Sum of numbers from 1 to 10")
	//Sum of first 10 natural numbers
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum of first 10 natural numbers:", sum)

	fmt.Println("Nested Loops Example")
	//Nested loops
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println("i:", i, "j:", j)
		}
	}
}
