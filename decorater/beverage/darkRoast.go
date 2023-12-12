package beverage

type DrakRoast struct{}

func (dr *DrakRoast) Cost() float64 {
	return 0.99
}

func (dr *DrakRoast) GetDescription() string {
	return "DrakRoast"
}
