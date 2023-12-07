package main

type ModelDuck struct {
	Duck
}

func GetModelDuck() *ModelDuck {
	md := new(ModelDuck)
	md.setFlyBehavior(&FlyRocketPowered{})
	md.setQuackBehavior(&Quack{})
	return md
}
