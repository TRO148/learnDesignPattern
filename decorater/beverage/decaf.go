package beverage

type Decaf struct{}

func (dr *Decaf) Cost() float64 {
	return 1.05
}

func (dr *Decaf) GetDescription() string {
	return "Decaf"
}
