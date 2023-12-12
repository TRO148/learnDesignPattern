package condiment

import "learn/decorater/beverage"

type Whip struct {
	Beverage beverage.Beverage
}

func (w *Whip) Cost() float64 {
	return w.Beverage.Cost() + 0.10
}

func (w *Whip) GetDescription() string {
	return w.Beverage.GetDescription() + ", Whip"
}
