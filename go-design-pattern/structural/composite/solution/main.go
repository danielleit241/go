package main

type Item interface {
	Cost() float64
}

type RealItem struct {
	Name  string
	Price float64
}

func (i RealItem) Cost() float64 {
	return i.Price
}

type Box struct {
	children []Item
}

func (b Box) Cost() float64 {
	var totalCost float64
	for _, child := range b.children {
		totalCost += child.Cost()
	}
	return totalCost
}

func CreatePackage() Item {
	return Box{
		children: []Item{
			RealItem{Name: "Mouse", Price: 20.5},
			Box{
				children: []Item{
					RealItem{Name: "Keyboard", Price: 60},
					RealItem{Name: "Charger", Price: 15},
				},
			},
		},
	}
}

func main() {
	packageItem := CreatePackage()
	println("Total cost of the package:", packageItem.Cost())
}
