package main

import "fmt"

type Duck interface {
	quack()
	fly()
}

type realDuck struct{}

func NewDuck() Duck {
	return &realDuck{}
}

func (d realDuck) quack() {
	fmt.Println("Duck Quack")
}

func (d realDuck) fly() {
	fmt.Println("Duck Fly")
}
