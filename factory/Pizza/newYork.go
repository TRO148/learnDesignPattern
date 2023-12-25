package pizza

import "factory/pizzaFactory"

type newYork struct {
	Pizza
}

func NewNewYork() *newYork {
	p := &newYork{Pizza: Pizza{Factory: &pizzaFactory.NewYork{}}}
	p.SetName("New York Style Pizza")
	return p
}

func (ny *newYork) Info() string {
	return ny.Pizza.Info()
}
