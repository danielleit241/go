package main

import (
	"fmt"

	"example.com/go/cat"
	"example.com/go/dog"
	"example.com/go/interfaces"
	"example.com/go/mouse"
)

func main() {
	dog, err := dog.New("Buddy")
	if err != nil {
		panic(err) // stop execution if error occurs
	}
	cat, err := cat.New("Whiskers")
	if err != nil {
		panic(err)
	}

	fmt.Print("Dog make sound plus: ")
	MakeSoundPlus(dog) // dog implements AnimalPlus
	fmt.Print("Cat make sound: ")
	MakeSound(cat) // cat implements Animal

	fmt.Print("Dog make action: ")
	MakeAction(dog) // dog implements AnimalActions
	fmt.Print("Cat make action: ")
	MakeAction(cat) // cat implements AnimalActions

	mouse, err := mouse.New("Mickey") // returns AnimalActions
	if err != nil {
		panic(err)
	}
	// fmt.Print(mouse.Eat() + "\n") // Error handled, interfaces.AnimalActions has no field or method Eat
	fmt.Print(mouse.Run() + "\n") // Directly call Run method
}

func MakeSound(a interfaces.Animal) {
	fmt.Println(a.Speak() + "\n")
	fmt.Println(a.Eat() + "\n")
}

func MakeAction(a interfaces.AnimalActions) {
	fmt.Println(a.Play() + "\n")
	fmt.Println(a.Run() + "\n")
}

func MakeSoundPlus(a interfaces.AnimalPlus) {
	fmt.Println(a.Speak() + "\n")
	fmt.Println(a.Eat() + "\n")
	fmt.Println(a.Extra() + "\n")
}
