package main

import (
	"factory/store"
	"fmt"
)

func main() {
	NewYorkPizza := store.NewYorkPizza()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>")
	ChicagoPizza := store.ChicagoPizza()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(NewYorkPizza.Info())
	fmt.Println(ChicagoPizza.Info())
}
