package interfaces

type Animal interface {
	Speak() string
	Eat() string
}

type AnimalPlus interface {
	Animal
	Extra() string
}

type AnimalActions interface {
	Play() string
	Run() string
}
