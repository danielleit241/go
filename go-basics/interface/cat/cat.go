package cat

import (
	"errors"
	"strings"
)

type Cat struct {
	Name string `json:"name"`
}

func New(name string) (*Cat, error) {
	name = strings.TrimSpace(name)
	if err := validateName(name); err != nil {
		return nil, err
	}
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

func (c *Cat) Eat() string {
	return c.Name + " is eating cat food."
}

func (c *Cat) Play() string {
	return c.Name + " is playing with a ball of yarn!"
}

func (c *Cat) Run() string {
	return c.Name + " is running swiftly!"
}
