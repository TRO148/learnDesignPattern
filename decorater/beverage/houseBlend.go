package beverage

type HouseBlend struct{}

func (dr *HouseBlend) Cost() float64 {
	return 0.89
}

func (dr *HouseBlend) GetDescription() string {
	return "HouseBlend"
}
