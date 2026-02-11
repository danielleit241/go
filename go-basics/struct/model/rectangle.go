package model

import "fmt"

type Rectangle struct {
	width  float64 //polimorphism via struct tags
	height float64
}

// factory function
// function package-level to create Rectangle instances
func NewRectangle(width, height float64) *Rectangle {
	if width <= 0 || height <= 0 {
		return nil
	}
	return &Rectangle{
		width:  width,
		height: height,
	}
}

// receiver functions (methods) for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r *Rectangle) PrintDetails() {
	fmt.Printf("Rectangle Details - Width: %.2f & Height: %.2f\n", r.width, r.height)
}
