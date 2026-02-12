package main

import "fmt"

func main() {
	// arrInt := []int{5, 2, 9, 1, 5, 6}
	// PrintArr(arrInt)

	// arrStr := []string{"banana", "apple", "orange"}
	// PrintArr(arrStr)

	var numbers [5]float64 // size 5 = slice 0:5
	numbers[2] = 10.0
	numbers[4] = 20.5
	// 	numbers[5] = 30.0 // Error: index out of range [5] with length 5
	fmt.Println(numbers)      // Print full array
	fmt.Println(numbers[1:])  // Slicing from index 1 to the end
	fmt.Println(numbers[4:5]) // Slicing from index 4 to 5
	fmt.Println(numbers[0:])  // Slicing from index 0 to the end
}

func PrintArr[T any](arr []T) {
	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %v\n", i, v)
	}
}
