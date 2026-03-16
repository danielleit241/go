package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type Point struct {
	X, Y int
}

func (p *Point) Clone() Prototype {
	return &Point{X: p.X, Y: p.Y}
}

type Node struct {
	Value    Point
	Children []Point
}

func (n *Node) Clone() Prototype {
	children := make([]Point, len(n.Children))

	for i, child := range n.Children {
		children[i] = *child.Clone().(*Point)
	}

	return &Node{
		Value:    *n.Value.Clone().(*Point),
		Children: children,
	}
}

type Color struct {
	red, green, blue int
}

func (c *Color) CloneWithRed(r int) Color {
	return Color{red: r, green: c.green, blue: c.blue} // Return a new Color object with the new red value
}

func main() {
	color := Color{100, 100, 100}
	color2 := color.CloneWithRed(200)
	fmt.Println(color, color2)

	p1 := Point{1, 2}
	p2 := p1.Clone()
	p1.X = 3
	fmt.Println(p1.X, p2.(*Point).X)

	n1 := Node{
		Value: p1,
		Children: []Point{
			{1, 2},
			{2, 3},
		},
	}
	n2 := n1.Clone().(*Node)
	n1.Children[0].X = 10
	fmt.Println(n1.Children[0].X, n2.Children[0].X)
}
