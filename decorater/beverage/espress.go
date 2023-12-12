package beverage

type Espress struct{}

func (dr *Espress) Cost() float64 {
	return 1.99
}

func (dr *Espress) GetDescription() string {
	return "Espress"
}
