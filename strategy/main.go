package main

func main() {
	mds := GetMiniDuckSimulator()
	mds.performFly()
	mds.performQuack()

	md := GetModelDuck()
	md.performFly()
	md.performQuack()
}
