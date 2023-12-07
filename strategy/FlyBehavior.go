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

type FlyRocketPowered struct{}

func (frp *FlyRocketPowered) fly() {
	println("i'm flying with a rocket!")
}
