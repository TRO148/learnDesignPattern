package condiment

import "learn/decorater/beverage"

type Milk struct {
	Beverage beverage.Beverage
}

func (m *Milk) Cost() float64 {
	return m.Beverage.Cost() + 0.10
}

func (m *Milk) GetDescription() string {
	return m.Beverage.GetDescription() + ", Milk"
}
