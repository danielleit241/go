package dog

import (
	"errors"
	"strings"
)

type Dog struct { //Name exported
	Name string `json:"name"`
}

func New(name string) (*Dog, error) {
	name = strings.TrimSpace(name)
	if err := validateName(name); err != nil {
		return nil, err
	}
	return &Dog{Name: name}, nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("dog name cannot be empty")
	}
	if len(name) < 2 {
		return errors.New("dog name must be at least 2 characters long")
	}
	if len(name) > 30 {
		return errors.New("dog name cannot exceed 30 characters")
	}
	return nil
}

func (d *Dog) Speak() string {
	return "Woof! I'm " + d.Name
}

func (d *Dog) Eat() string {
	return d.Name + " is eating dog food."
}

func (d *Dog) Extra() string {
	return d.Name + " loves to play fetch!"
}

func (d *Dog) Play() string {
	return d.Name + " is playing happily!"
}

func (d *Dog) Run() string {
	return d.Name + " is running fast!"
}

func (d *Dog) String() string { // Implementing the Stringer interface for better string representation of Dog
	return "Name: " + d.Name + "\n Speak: " + d.Speak() + "\n Eat: " + d.Eat() + "\n Extra: " + d.Extra() + "\n Play: " + d.Play() + "\n Run: " + d.Run()
}
