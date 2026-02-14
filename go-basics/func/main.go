package main

import (
	"fmt"
	"time"
)

//closure function

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	//menu()

	// var number1, number2 int
	// var op string
	// fmt.Print("Enter first number: ")
	// fmt.Scan(&number1)
	// fmt.Print("Enter second number: ")
	// fmt.Scan(&number2)
	// fmt.Print("Enter operation (+, -, *, /): ")
	// fmt.Scan(&op)
	// result := mathOperation(number1, number2, op)
	// fmt.Printf("The result is: %f\n", result)
	// printOperation(number1, number2, op)

	// year := time.Now().Year()
	// printLeapYear(year)

	// n := 10
	// fmt.Printf("Fibonacci of %d is %d\n", n, fibonacci(n))
	// countdown(5)

	// a, b := "hello", "world"
	// fmt.Printf("Before swap: a = %s, b = %s\n", a, b)
	// a, b = swap(a, b)
	// fmt.Printf("After swap: a = %s, b = %s\n", a, b)

	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println("Positive:", pos(i), "Negative:", neg(-2*i))
	}
}

// return multiple values function
func swap(a, b string) (string, string) {
	return b, a
}

// recursive function
func countdown(n int) {
	if n <= 0 {
		fmt.Println("Countdown finished!")
		return
	}
	fmt.Println(n)
	time.Sleep(200 * time.Millisecond)
	countdown(n - 1)
}

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

// void function
func printOperation(number1, number2 int, op string) {
	switch op {
	case "+":
		fmt.Printf("%d + %d = %d\n", number1, number2, number1+number2)
	case "-":
		fmt.Printf("%d - %d = %d\n", number1, number2, number1-number2)
	case "*":
		fmt.Printf("%d * %d = %d\n", number1, number2, number1*number2)
	case "/":
		if number2 != 0 {
			fmt.Printf("%d / %d = %f\n", number1, number2, float64(number1)/float64(number2))
		}
	default:
		fmt.Println("Invalid operation")
	}
}

func printLeapYear(year int) {
	if isLeapYear(year) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}

// return function
func mathOperation(number1, number2 int, op string) float64 {
	switch op {
	case "+":
		return float64(number1 + number2)
	case "-":
		return float64(number1 - number2)
	case "*":
		return float64(number1 * number2)
	case "/":
		if number2 != 0 {
			return float64(number1) / float64(number2)
		}
	}
	return 0
}

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}
