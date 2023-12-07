package main

type QuackBehavior interface {
	quack()
}

type Quack struct{}

func (q *Quack) quack() {
	println("Quack")
}

type MuteQuack struct{}

func (mq *MuteQuack) quack() {
	println("<< Silence >>")
}

type Squeak struct{}

func (s *Squeak) quack() {
	println("Squeak")
}
