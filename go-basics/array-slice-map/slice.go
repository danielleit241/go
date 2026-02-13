package main

import (
	"fmt"
	"reflect"
	"slices"
)

func SliceExample() {
	IntroductionSlice()

	SliceInitialization()

	TraverseSlice()

	AppendSlice()

	SubSlice()

	SlicesPackage()
}

func IntroductionSlice() {
	fmt.Println("Slice Introduction Examples")

	//Slice declaration and initialization

	numbers := []float64{10.0, 20.5, 30.0, 40.5, 50.0}

	fmt.Println(numbers)

	// Skice declaration and initialization -> does not have a fixed size
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Array declaration and initialization -> has a fixed size
	array := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array)

	fmt.Println("How to know which is array and which is slice?")
	fmt.Println("Using reflect.TypeOf")
	fmt.Println("Array:", reflect.TypeOf(array).Kind())
	fmt.Println("Slice:", reflect.TypeOf(slice).Kind())
}

func SliceInitialization() {
	fmt.Println("Slice Initialization Examples")

	array := [5]string{"Go", "Python", "Java", "C++", "JavaScript"}
	fmt.Println("Array:", array)
	fmt.Println("Is array an Array?", reflect.TypeOf(array).Kind() == reflect.Array)

	slice0 := array[1:4] // Slicing the array from index 1 to 3
	fmt.Println("Slice:", slice0)
	fmt.Println("Is slice0 a slice?", reflect.TypeOf(slice0).Kind() == reflect.Slice)

	slice1 := make([]int, 3, 5) // Creating a slice of integers with length 3 and capacity 5 with default values of data type
	var slice2 []int
	fmt.Println("Slice1:", slice1)             // [0 0 0 _ _]
	fmt.Println("Cap of slice1:", cap(slice1)) // 5
	fmt.Println("Slice2:", slice2)             // []
	fmt.Println("Cap of slice2:", cap(slice2)) // 0

	slice1 = append(slice1, 10, 20, 30)
	fmt.Println("Slice1 after appending elements:", slice1) // [0 0 0 10 20 30]
	fmt.Println("Cap of slice1:", cap(slice1))              // 10 -> capacity doubles when exceeded
}

func TraverseSlice() {
	fmt.Println("Traverse Slice Examples")

	apples := []string{"Red Delicious", "Granny Smith", "Fuji", "Gala", "Honeycrisp"}

	for i := 0; i < len(apples); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, apples[i])
	}

	for index, value := range apples {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	matrix := [][]int{ // 2D Slice
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}

	for i, row := range matrix {
		for j, value := range row {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, value)
		}
	}
}

func AppendSlice() {
	fmt.Println("Append Slice Examples")

	slice1 := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice1:", slice1)
	slice2 := []int{60, 70, 80, 90, 100}
	fmt.Println("Slice2:", slice2)

	slice1 = append(slice1, 36)
	fmt.Println("After appending 36 to slice1:", slice1)
	slice1 = append(slice1, slice2...)
	fmt.Println("After appending slice2 to slice1:", slice1)
}

// Slice does not create a new array when slicing, it creates a new slice header that points to the same underlying array. Therefore, modifying the subslice will affect the original slice and vice versa. However, when appending to a subslice exceeds its capacity, a new underlying array is created, and the original slice remains unchanged. This is because the subslice's capacity is determined by the original slice's length from the starting index of the subslice.

// Tips:
// 1. SubSlices do not have memory for therselves
// 2. SubSlices are a view into the original slice, they share the same underlying array
// 3. When a subslice overflows its capacity, it creates a new underlying array and the original slice remains unchanged
func SubSlice() {
	fmt.Println("SubSlice Examples")
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	subSliceNumbers1 := numbers[2:7]
	subSliceNumbers2 := numbers[0:]
	subSliceNumbers3 := numbers[:4]
	subSliceNumbers4 := numbers[4:8]
	fmt.Println("Original Slice:", numbers, "length:", len(numbers), "capacity:", cap(numbers))
	fmt.Println("SubSlice1 (index 2 to 6):", subSliceNumbers1, "length:", len(subSliceNumbers1), "capacity:", cap(subSliceNumbers1))
	fmt.Println("SubSlice2 (index 0 to end):", subSliceNumbers2, "length:", len(subSliceNumbers2), "capacity:", cap(subSliceNumbers2))
	fmt.Println("SubSlice3 (start to index 3):", subSliceNumbers3, "length:", len(subSliceNumbers3), "capacity:", cap(subSliceNumbers3))
	fmt.Println("SubSlice4 (index 4 to 7):", subSliceNumbers4, "length:", len(subSliceNumbers4), "capacity:", cap(subSliceNumbers4))

	// Modifying subslice affects the original slice
	subSliceNumbers1[0] = 999
	subSliceNumbers2[3] = 888
	subSliceNumbers1 = append(subSliceNumbers1, 777, 666, 555, 444)
	fmt.Println("After modifying SubSlice1:")
	fmt.Println("Original Slice:", numbers, "length:", len(numbers), "capacity:", cap(numbers))
	fmt.Println("SubSlice1:", subSliceNumbers1, "length:", len(subSliceNumbers1), "capacity:", cap(subSliceNumbers1))
}

func SlicesPackage() {
	fmt.Println("Slices Package Examples")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original slice:", slice)

	copied := slices.Clone(slice)
	fmt.Println("Copied slice:", copied)

	compareSlice := slices.Equal(slice, copied)
	fmt.Println("Are slices equal?", compareSlice)

	findPos := slices.Index(slice, 11) // returns the index of the first occurrence of 11, or -1 if not found
	fmt.Println("Position of 11 in slice:", findPos)

	slices.Reverse(slice)
	fmt.Println("Reversed slice:", slice)

	itemIsExist := slices.Contains(slice, 5)
	fmt.Println("Does slice contain 5?", itemIsExist)

	sortedSlice := []int{3, 1, 4, 5, 2}
	slices.Sort(sortedSlice)
	fmt.Println("Sorted slice:", sortedSlice)

	insertedSlice := slices.Insert(sortedSlice, 2, 10)
	fmt.Println("Slice after inserting 10 at index 2:", insertedSlice)

	removedSlice := slices.Delete(insertedSlice, 4, 6) // removes elements from index 4 to 5
	fmt.Println("Slice after deleting elements from index 4 to 5:", removedSlice)

	max := slices.Max(removedSlice)
	min := slices.Min(removedSlice)
	fmt.Println("Max value in slice:", max)
	fmt.Println("Min value in slice:", min)
}
