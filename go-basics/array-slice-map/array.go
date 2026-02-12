package main

import (
	"fmt"

	"example.com/go/employee"
)

func ArrayExample() {
	// arrInt := []int{5, 2, 9, 1, 5, 6}
	// PrintArr(arrInt)

	// arrStr := []string{"banana", "apple", "orange"}
	// PrintArr(arrStr)

	// var numbers [5]float64
	// numbers[2] = 10.0
	// numbers[4] = 20.5
	// // 	numbers[5] = 30.0 // Error: index out of range [5] with length 5
	// fmt.Println(numbers)      // Print full array
	// fmt.Println(numbers[1:])  // Slicing from index 1 to the end
	// fmt.Println(numbers[4:5]) // Slicing from index 4 to 5
	// fmt.Println(numbers[0:])  // Slicing from index 0 to the end

	// var matrix [3][3]int
	// matrix[0] = [3]int{1, 2, 3}
	// matrix[1] = [3]int{4, 5, 6}
	// matrix[2] = [3]int{7, 8, 9}
	// fmt.Println(matrix)

	numbers := []float64{10.0, 20.5, 30.0, 40.5, 50.0}
	printArrUsingForIndex(numbers)
	printArrUsingForRange(numbers)

	employees := [...]*employee.Employee{
		employee.NewEmployee(1, "John Doe", "Developer", 60000.0),
		employee.NewEmployee(2, "Jane Smith", "Designer", 55000.0),
	}

	printArrUsingForRange(employees[:])
}

func printArrUsingForIndex[T any](arr []T) {
	fmt.Println("Printing array using for loop with index:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Index: %d, Value: %v\n", i, arr[i])
	}
}

func printArrUsingForRange[T any](arr []T) {
	fmt.Println("Printing array using for loop with range:")
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %v\n", index, value)
	}
}
