package main

import (
	"fmt"

	"example.com/go/cat"
	"example.com/go/dog"
	"example.com/go/interfaces"
)

func main() {
	dog := dog.New("Buddy")
	cat := cat.New("Whiskers")

	fmt.Print("Dog speak: ")
	MakeSound(dog)

	fmt.Println()

	fmt.Print("Cat speak: ")
	MakeSound(cat)
}

func MakeSound(a interfaces.Animal) {
	fmt.Println(a.Speak())
}
