package condiment

import "learn/decorater/beverage"

type Soy struct {
	Beverage beverage.Beverage
}

func (s *Soy) Cost() float64 {
	return s.Beverage.Cost() + 0.15
}

func (s *Soy) GetDescription() string {
	return s.Beverage.GetDescription() + ", Soy"
}
