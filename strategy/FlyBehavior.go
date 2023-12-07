package main

type FlyBehavior interface {
	fly()
}

type FlyWithWings struct{}

func (f *FlyWithWings) fly() {
	println("I'm flying!")
}

type FlyNoWay struct{}

func (f *FlyNoWay) fly() {
	println("I can't fly")
}
