package pizzaFactory

import (
	"factory/cheese"
	"factory/clams"
	"factory/dough"
	"factory/sauce"
)

type Chicago struct {
}

func (ny *Chicago) CreateDough() dough.Interface {
	return &dough.ThinCrust{}
}

func (ny *Chicago) CreateSauce() sauce.Interface {
	return &sauce.Marinara{}
}

func (ny *Chicago) CreateCheese() cheese.Interface {
	return &cheese.Reggiano{}
}

func (ny *Chicago) CreateClam() clams.Interface {
	return &clams.Fresh{}
}
