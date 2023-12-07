package main

type MiniDuckSimulator struct {
	Duck
}

func GetMiniDuckSimulator() *MiniDuckSimulator {
	mds := new(MiniDuckSimulator)
	mds.setFlyBehavior(&FlyWithWings{})
	mds.setQuackBehavior(&Quack{})
	return mds
}
