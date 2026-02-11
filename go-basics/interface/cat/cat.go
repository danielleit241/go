package cat

type Cat struct {
	Name string `json:"name"`
}

func New(name string) *Cat {
	if name == "" {
		name = "MeoMeo"
	}
	return &Cat{Name: name}
}

func (c Cat) Speak() string {
	return "Meow! I'm " + c.Name
}
