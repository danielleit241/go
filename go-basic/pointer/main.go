package main

import "fmt"

func main() {
	// var a, b int = 10, 20
	// fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	// swap(&a, &b)
	// fmt.Printf("After swap: a = %d, b = %d\n", a, b)

	// x := 42
	// printAddress(x)

	// pointerX := &x
	// fmt.Printf("Address of x in main: %p\n", pointerX)
	// fmt.Printf("Value of x using pointer: %d\n", *pointerX)

	// word := "Original"
	// fmt.Printf("Before change: %s\n", word)
	// changeWordUsingPointer(&word)
	// fmt.Printf("After change: %s\n", word)

	// passByReference()

	// num := 25
	// fmt.Printf("Before function call: num = %d, address = %p\n", num, &num)
	// passByValue(num)

	// multilevelPointer()
}

func multilevelPointer() {
	var value = 100
	ptr1 := &value
	ptr2 := &ptr1
	ptr3 := &ptr2

	fmt.Printf("Value: %d\n", value)
	fmt.Printf("Value using ptr1: %d with address %p\n", *ptr1, ptr1)
	fmt.Printf("Value using ptr2: %d with address %p\n", **ptr2, ptr2)
	fmt.Printf("Value using ptr3: %d with address %p\n", ***ptr3, ptr3)
}

func passByValue(num int) {
	fmt.Printf("After function call: num = %d, address = %p\n", num, &num)
}

func passByReference() { //name is a pointer to string
	name := "Danile"
	ptrName := &name // pointer to name variable

	ptrName2 := ptrName // another pointer to the same variable

	fmt.Printf("Find value using pointer: %v\n", *ptrName)
	fmt.Printf("Address of name variable: %p\n", ptrName)

	changeWordUsingPointer(&name)

	fmt.Printf("Name after change using pointer: %s\n", name)
	fmt.Printf("Address of name variable after change: 	%p\n", ptrName)
	fmt.Printf("Find value using another pointer: %v\n", *ptrName2)
	fmt.Printf("Address using another pointer: %p\n", ptrName2)
}

func changeWordUsingPointer(wordPtr *string) {
	*wordPtr = "Changed via pointer"
}

func printAddress(x int) {
	fmt.Printf("Address of x inside function: %p\n", &x)
	fmt.Printf("Value of x inside function: %d\n", x)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
