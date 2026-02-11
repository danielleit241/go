package interfaces

// If an interface has only one method, it is called a functional interface.
// Name convention is to use -er suffix.
type Speaker interface {
	Speak() string
}
