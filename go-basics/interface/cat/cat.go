package cat

import (
	"errors"
	"strings"
)

type Cat struct {
	Name string `json:"name"`
}

func New(name string) (*Cat, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}
	name = strings.TrimSpace(name)
	return &Cat{Name: name}, nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("cat name cannot be empty")
	}
	if len(name) < 2 {
		return errors.New("cat name must be at least 2 characters long")
	}
	if len(name) > 30 {
		return errors.New("cat name cannot exceed 30 characters")
	}
	return nil
}

func (c *Cat) Speak() string {
	return "Meow! I'm " + c.Name
}
