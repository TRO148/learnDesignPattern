package pizza

import (
	"factory/cheese"
	"factory/clams"
	"factory/dough"
	"factory/pizzaFactory"
	"factory/sauce"
	"fmt"
)

type Interface interface {
	Info() string
	Prepare()
	Bake()
	Cut()
	Box()
}

type Pizza struct {
	Factory pizzaFactory.Interface
	name    string

	sauce  sauce.Interface
	dough  dough.Interface
	cheese cheese.Interface
	clam   clams.Interface
}

func (p *Pizza) Prepare() {
	p.sauce = p.Factory.CreateSauce()
	p.dough = p.Factory.CreateDough()
	p.cheese = p.Factory.CreateCheese()
	p.clam = p.Factory.CreateClam()

	fmt.Println("Preparing " + p.name)
}

func (p *Pizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (p *Pizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p *Pizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (p *Pizza) SetName(name string) {
	p.name = name
}

func (p *Pizza) Info() string {
	return fmt.Sprintf("Pizza name: %s\nSauce: %s\nDough: %s\nCheese: %s\nClam: %s\n", p.name, p.sauce.Get(), p.dough.Get(), p.cheese.Get(), p.clam.Get())
}
