package pizza

import "factory/pizzaFactory"

type chicago struct {
	Pizza
}

func NewChicago() *chicago {
	p := &chicago{Pizza: Pizza{Factory: &pizzaFactory.Chicago{}}}
	p.SetName("Chicago Style Pizza")
	return p
}

func (ny *chicago) Info() string {
	return ny.Pizza.Info()
}
