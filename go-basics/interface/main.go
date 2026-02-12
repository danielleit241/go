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

	fmt.Println("Empty interface: ")
	PrintValueV1("The dog name is: " + dog.Name + "\n")
	PrintValueV2("The dog name is: " + dog.Name)
	PrintValues(42)
	PrintValues(true)
	PrintStringValue("Hello, World!")
	// PrintStringValue(100) // Not a string value
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

// PrintName demonstrates type assertion with an empty interface
func PrintValueV1(value interface{}) {
	fmt.Print(value)
}

// PrintNameV2 demonstrates type assertion with an empty interface using 'any' keyword = 'interface{}'
func PrintValueV2(value any) {
	fmt.Print(value)
}

// any = interface{} = any type
// Because any can hold any type, we need to use `type assertion` or type switch to get the underlying value and type.

func PrintValues(value any) {
	switch v := value.(type) {
	case string:
		fmt.Println("String value:", v)
	case int:
		fmt.Println("Integer value:", v)
	case bool:
		fmt.Println("Boolean value:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func PrintStringValue(value any) {
	strValue, ok := value.(string)
	if !ok {
		fmt.Println("Not a string value")
		return
	}
	fmt.Println("String value:", strValue)
}
