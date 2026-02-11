package main

import (
	"fmt"

	"example.com/go/cat"
	"example.com/go/dog"
	"example.com/go/interfaces"
)

func main() {
	dog, err := dog.New("Buddy")
	if err != nil {
		fmt.Println("Error creating dog:", err)
		return
	}
	cat, err := cat.New("Whiskers")
	if err != nil {
		fmt.Println("Error creating cat:", err)
		return
	}

	fmt.Print("Dog speak: ")
	MakeSound(dog)

	fmt.Println()

	fmt.Print("Cat speak: ")
	MakeSound(cat)
}

func MakeSound(a interfaces.Animal) {
	fmt.Println(a.Speak())
}
