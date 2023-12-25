package pizzaFactory

import (
	"factory/cheese"
	"factory/clams"
	"factory/dough"
	"factory/sauce"
)

type Interface interface {
	CreateDough() dough.Interface
	CreateSauce() sauce.Interface
	CreateCheese() cheese.Interface
	CreateClam() clams.Interface
}
