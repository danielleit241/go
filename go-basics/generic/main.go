package main

import (
	"cmp"
	"fmt"
)

// Generic in Go is primarily implemented through the use of type parameters, which allow you to write functions, types, and methods that can operate on different data types without sacrificing type safety. Generics were introduced in Go 1.18.

func main() {
	PrintValue("Hello, World!")
	PrintValue(42)

	if AreEqual(10, 10) {
		fmt.Println("10 and 10 are equal")
	}

	if !AreEqual("Hello", "hello") {
		fmt.Println("Hello and hello are not equal")
	}

	maxInt := Max(10.2, 20)
	fmt.Println("Max of 10.2 and 20 is:", maxInt)

	maxStr := Max("apple", "banana")
	fmt.Println("Max of apple and banana is:", maxStr)
}

func PrintValue[T any](value T) {
	fmt.Println("Value:", value)
}

// Comparable constraint allows types that support comparison operators (==, !=)
// such as integers, floats, strings, pointers, and structs/arrays composed of comparable types. Not slices, maps, or functions.
func AreEqual[T comparable](a, b T) bool {
	return a == b
}

func AreNotEqual[T comparable](a, b T) bool {
	return a != b
}

// Using the Ordered constraint from the cmp package to allow types that support ordering operators (<, <=, >, >=)

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
