package dog

type Dog struct { //Name exported
	Name string `json:"name"`
}

func New(name string) *Dog {
	if name == "" {
		name = "Kiki"
	}
	return &Dog{Name: name}
}

func (d Dog) Speak() string {
	return "Woof! I'm " + d.Name
}
