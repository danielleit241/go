package main

import "fmt"

// Generic in Go is primarily implemented through the use of type parameters, which allow you to write functions, types, and methods that can operate on different data types without sacrificing type safety. Generics were introduced in Go 1.18.

func main() {
	PrintValue("Hello, World!")
	PrintValue(42)
}

func PrintValue[T any](value T) {
	fmt.Println("Value:", value)
}
