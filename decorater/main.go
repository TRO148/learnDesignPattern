package main

import (
	"fmt"
	"learn/decorater/beverage"
	"learn/decorater/condiment"
	"learn/decorater/model"
)

func main() {
	var beverage1 beverage.Beverage = &beverage.Espress{}
	beverage1 = &condiment.Mocha{Beverage: beverage1}
	beverage1 = &condiment.Milk{Beverage: beverage1}
	beverage1 = &condiment.Milk{Beverage: beverage1}

	tallResult := model.Tall(beverage1)
	fmt.Println(tallResult.GetDescription())
	fmt.Println(tallResult.Cost())

	var beverage2 beverage.Beverage = &beverage.HouseBlend{}
	beverage2 = &condiment.Milk{Beverage: beverage2}
	beverage2 = &condiment.Soy{Beverage: beverage2}
	beverage2 = &condiment.Whip{Beverage: beverage2}

	grandeResult := model.Grande(beverage2)
	fmt.Println(grandeResult.GetDescription())
	fmt.Printf("%0.2f", grandeResult.Cost())
}
