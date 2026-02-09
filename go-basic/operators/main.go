package main

import "fmt"

func main() {
	fmt.Println("Math Operators:")
	mathOperations()

	fmt.Println("\nAssign Operators:")
	assignOperators()

	fmt.Println("\nComparison Operators:")
	comparisonOperators()

	fmt.Println("\nLogical Operators:")
	logicalOperators()
}

func logicalOperators() {
	a := true
	b := false
	fmt.Printf("%v AND %v: %t\n", a, b, a && b)
	fmt.Printf("%v OR %v: %t\n", a, b, a || b)
	fmt.Printf("NOT %v: %t\n", a, !a)
}

func comparisonOperators() {
	a := 10
	b := 20
	fmt.Printf("Is %d equal to %d? %t\n", a, b, a == b)
	fmt.Printf("Is %d not equal to %d? %t\n", a, b, a != b)
	fmt.Printf("Is %d greater than %d? %t\n", a, b, a > b)
	fmt.Printf("Is %d less than %d? %t\n", a, b, a < b)
	fmt.Printf("Is %d greater than or equal to %d? %t\n", a, b, a >= b)
	fmt.Printf("Is %d less than or equal to %d? %t\n", a, b, a <= b)
}

func assignOperators() {
	var x int
	var y int
	x = 10
	y = 20
	fmt.Printf("x: %d, y: %d\n", x, y)

	x += 5
	fmt.Printf("After x += 5, x: %d\n", x)

	y *= 2
	fmt.Printf("After y *= 2, y: %d\n", y)

	x -= 3
	fmt.Printf("After x -= 3, x: %d\n", x)
}

func mathOperations() {
	a := 13
	b := 5
	i := 0
	j := 0

	// Addition
	sum := a + b
	fmt.Printf("The result of %d + %d is: %d\n", a, b, sum)

	// Subtraction
	diff := a - b
	fmt.Printf("The result of %d - %d is: %d\n", a, b, diff)

	// Multiplication
	prod := a * b
	fmt.Printf("The result of %d * %d is: %d\n", a, b, prod)

	// Division
	quotient := float64(a) / float64(b)
	fmt.Printf("The result of %d / %d is: %f\n", a, b, quotient)

	// Modulus
	mod := a % b
	fmt.Printf("The result of %d %% %d is: %d\n", a, b, mod)

	// Increment
	i++
	fmt.Printf("After incrementing, i is: %d\n", i)

	// Decrement
	j--
	fmt.Printf("After decrementing, j is: %d\n", j)
}
