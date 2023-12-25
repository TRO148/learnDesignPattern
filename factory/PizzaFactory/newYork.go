package pizzaFactory

import (
	"factory/cheese"
	"factory/clams"
	"factory/dough"
	"factory/sauce"
)

type NewYork struct {
}

func (ny *NewYork) CreateDough() dough.Interface {
	return &dough.ThickCrust{}
}

func (ny *NewYork) CreateSauce() sauce.Interface {
	return &sauce.PlumTomato{}
}

func (ny *NewYork) CreateCheese() cheese.Interface {
	return &cheese.Mozzarella{}
}

func (ny *NewYork) CreateClam() clams.Interface {
	return &clams.Frozen{}
}
