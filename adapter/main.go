package main

import "fmt"

func main() {
	var duck Duck
	duck = NewDuck()
	var turkey Turkey
	turkey = newTurkey()

	fmt.Printf("%T\n", duck)
	duck.quack()
	fmt.Printf("%T\n", turkey)
	turkey.gobble()

	duck = NewAdapter(turkey)
	duck.quack()

	turkey = NewAdapter(duck)
	turkey.gobble()

	fmt.Printf("%T\n", duck)
	duck.fly()
	fmt.Printf("%T\n", turkey)
	turkey.fly()
}
