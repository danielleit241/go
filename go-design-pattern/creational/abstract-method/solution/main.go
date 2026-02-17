package problem

import (
	"fmt"
)

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

// SOLUTION:

type VoucherAbstractFactory interface {
	GetDrink() Drink
	GetFood() Food
}

type CoffeMorningVoucherFactory struct{}

func (c *CoffeMorningVoucherFactory) GetDrink() Drink {
	return &Coffee{}
}

func (c *CoffeMorningVoucherFactory) GetFood() Food {
	return &Cake{}
}

type BeerEveningVoucherFactory struct{}

func (b *BeerEveningVoucherFactory) GetDrink() Drink {
	return &Beer{}
}

func (b *BeerEveningVoucherFactory) GetFood() Food {
	return &GrilledOctopus{}
}

// Using Factory Method pattern to create vouchers, we can ensure that only valid combinations of drink and food are created, adhering to the Open/Closed Principle.

func GetVoucherFactory(campaignName string) (VoucherAbstractFactory, error) {
	if campaignName == "morning" {
		return &CoffeMorningVoucherFactory{}, nil
	}
	if campaignName == "evening" {
		return &BeerEveningVoucherFactory{}, nil
	}
	return nil, fmt.Errorf("unknown campaign name: %s", campaignName)
}

func GetVoucher(factory VoucherAbstractFactory) Voucher {
	return Voucher{
		Drink: factory.GetDrink(),
		Food:  factory.GetFood(),
	}
}

func main() {
	voucherFac, err := GetVoucherFactory("morning")
	if err != nil {
		panic(err)
	}
	voucher := GetVoucher(voucherFac)
	fmt.Printf("%+v\n", voucher)
}
