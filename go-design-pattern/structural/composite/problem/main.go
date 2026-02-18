package problem

type Item struct {
	Name  string
	Price float64

	children []*Item
}

func (i *Item) Cost() float64 {
	cost := i.Price
	for _, child := range i.children {
		cost += child.Cost()
	}
	return cost
}

func CreatePackage() *Item {
	return &Item{
		// It's a box not an item
		Name: "root box",
		// So the price is 0
		Price: 0,
		// But it has children
		children: []*Item{
			// Here is an item
			{Name: "Mouse", Price: 20.5, children: nil},
			// Here is another box contains items
			{Name: "sub box", Price: 0, children: []*Item{
				{Name: "Keyboard", Price: 60, children: nil},
				{Name: "Charger", Price: 15, children: nil},
			}},
		},
	}
}

func main() {
	packageItem := CreatePackage()
	println("Total cost of the package:", packageItem.Cost())
}
