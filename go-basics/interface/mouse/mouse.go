package mouse

import (
	"errors"
	"strings"

	"github.com/danielleit241/interfaces"
)

type Mouse struct {
	Name string `json:"name"`
}

// When creating a new Mouse, we return the interface type AnimalActions so that only those methods are accessible.
func New(name string) (interfaces.AnimalActions, error) { // return type changed to AnimalActions
	name = strings.TrimSpace(name)
	if err := validateName(name); err != nil {
		return nil, err
	}
	return &Mouse{Name: name}, nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("mouse name cannot be empty")
	}
	if len(name) < 2 {
		return errors.New("mouse name must be at least 2 characters long")
	}
	if len(name) > 30 {
		return errors.New("mouse name cannot exceed 30 characters")
	}
	return nil
}

func (m *Mouse) Speak() string {
	return "Squeak! I'm " + m.Name
}

func (m *Mouse) Eat() string {
	return m.Name + " is eating cheese."
}

func (m *Mouse) Play() string {
	return m.Name + " is playing with a tiny ball!"
}

func (m *Mouse) Run() string {
	return m.Name + " is running around quickly!"
}
