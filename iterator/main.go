package main

import "container/list"

func main() {
	dinerMenu := &DinerMenu{menuItems: []Item{"Vegetarian BLT", "BLT", "Soup of the day", "Hotdog", "Steamed Veggies and Brown Rice", "Pasta"}}

	l := list.New()
	l.PushBack(Item("K&B's Pancake Breakfast"))
	l.PushBack(Item("Regular Pancake Breakfast"))
	pancakeHouseMenu := &PancakeHouseMenu{menuItems: *l}

	cafeMenu := &CafeMenu{menuItems: map[string]Item{"Veggie Burger and Air Fries": "Veggie Burger and Air Fries", "Soup of the day": "Soup of the day", "Burrito": "Burrito"}}

	waitress := NewWaitress(dinerMenu, pancakeHouseMenu, cafeMenu)
	waitress.PrintMenu()
}
