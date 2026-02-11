package dog

import (
	"errors"
	"strings"
)

type Dog struct { //Name exported
	Name string `json:"name"`
}

func New(name string) (*Dog, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}
	name = strings.TrimSpace(name)
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
