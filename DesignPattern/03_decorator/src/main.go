package main

import "fmt"

type Beverage interface {
	GetDescription() string
	GetCost() float64
}

type Coffee struct {
	desc string
	cost float64
}

func NewCoffee(desc string, cost float64) *Coffee {
	return &Coffee{desc, cost}
}

func (c *Coffee) GetDescription() string {
	return c.desc
}

func (c *Coffee) GetCost() float64 {
	return c.cost
}

type Decorator struct {
	b Beverage
	Coffee
}

func NewDecorator(desc string, cost float64, b Beverage) *Decorator {
	return &Decorator{
		b:      b,
		Coffee: Coffee{desc: desc, cost: cost},
	}
}

func (d *Decorator) GetDescription() string {
	return d.b.GetDescription() + ", " + d.desc
}

func (d *Decorator) GetCost() float64 {
	return d.b.GetCost() + d.cost
}

func main() {
	b1 := NewDecorator("milk", .1, NewDecorator("salt", .2, NewCoffee("coffee1", 1000)))
	fmt.Printf("b1:%#v cost: $%#v\n", b1.GetDescription(), b1.GetCost())

	var b2 Beverage = NewCoffee("coffee2", 2000)
	b2 = NewDecorator("milk", .2, b2)
	b2 = NewDecorator("milk", .2, b2)
	b2 = NewDecorator("salt", 1.0, b2)
	fmt.Printf("b2:%#v cost: $%#v\n", b2.GetDescription(), b2.GetCost())
}
