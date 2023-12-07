package main

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (duck *Duck) setFlyBehavior(fb FlyBehavior) {
	duck.flyBehavior = fb
}

func (duck *Duck) setQuackBehavior(qb QuackBehavior) {
	duck.quackBehavior = qb
}

func (duck *Duck) performFly() {
	duck.flyBehavior.fly()
}

func (duck *Duck) performQuack() {
	duck.quackBehavior.quack()
}

func (duck *Duck) swim() {
	println("All ducks float, even decoys")
}
