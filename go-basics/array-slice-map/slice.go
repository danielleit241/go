package main

import (
	"fmt"
	"reflect"
)

func SliceExample() {
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
