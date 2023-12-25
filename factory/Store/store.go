package store

import pizza "factory/Pizza"

func makePizza(p pizza.Interface) pizza.Interface {
	p.Prepare()
	p.Bake()
	p.Cut()
	p.Box()
	return p
}

func NewYorkPizza() pizza.Interface {
	NewYorkPizza := pizza.NewNewYork()
	return makePizza(NewYorkPizza)
}

func ChicagoPizza() pizza.Interface {
	ChicagoPizza := pizza.NewChicago()
	return makePizza(ChicagoPizza)
}
