package condiment

import "learn/decorater/beverage"

type Mocha struct {
	Beverage beverage.Beverage
}

func (m *Mocha) Cost() float64 {
	return m.Beverage.Cost() + 0.20
}

func (m *Mocha) GetDescription() string {
	return m.Beverage.GetDescription() + ", Mocha"
}
