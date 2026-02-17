package problem

import "fmt"

type Drink interface {
	Drink()
}

type Food interface {
	Eat()
}

type Voucher struct {
	Drink
	Food
}

type Coffee struct{}

func (c *Coffee) Drink() {
	println("It's a coffee, drinkable")
}

type Beer struct{}

func (b *Beer) Drink() {
	println("It's a beer, drinkable")
}

type Cake struct{}

func (c *Cake) Eat() {
	println("It's a cake, eatable")
}

type GrilledOctopus struct{}

func (g *GrilledOctopus) Eat() {
	println("It's a grilled octopus, eatable")
}

func main() {
	fmt.Println([]Voucher{
		// Coffee and cake are a good
		Voucher{
			Drink: &Coffee{},
			Food:  &Cake{},
		},

		// Beer and grilled octopus are a good combination
		Voucher{
			Drink: &Beer{},
			Food:  &GrilledOctopus{},
		},

		// Coffee and grilled octopus are not a good combination, can we create a voucher for this combination?
		// For coding simplicity, we can create a voucher for this combination, but it doesn't make sense in real life.
		Voucher{
			Drink: &Coffee{},
			Food:  &GrilledOctopus{},
		},
	})
}
