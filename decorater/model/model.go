package model

import (
	"learn/decorater/beverage"
	"sort"
	"strings"
)

type BaseBeverage struct {
	cost        float64
	description string
}

func (bb *BaseBeverage) Cost() float64          { return bb.cost }
func (bb *BaseBeverage) GetDescription() string { return bb.description }

func Tall(b beverage.Beverage) beverage.Beverage {
	condiments := b.GetDescription() + ", Tall"
	cost := b.Cost() + 0.10

	newBeverage := &BaseBeverage{cost: cost, description: formatString(condiments)}
	return newBeverage
}

func Grande(b beverage.Beverage) beverage.Beverage {
	condiments := b.GetDescription() + ", Grande"
	cost := b.Cost() + 0.15

	newBeverage := &BaseBeverage{cost: cost, description: formatString(condiments)}
	return newBeverage
}

func Venti(b beverage.Beverage) beverage.Beverage {
	condiments := b.GetDescription() + ", Venti"
	cost := b.Cost() + 0.20

	newBeverage := &BaseBeverage{cost: cost, description: formatString(condiments)}
	return newBeverage
}

func formatString(s string) string {
	words := strings.Split(s, ", ")
	sort.Strings(words)
	return strings.Join(words, ", ")
}
