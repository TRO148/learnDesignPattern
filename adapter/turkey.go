package main

import "fmt"

type Turkey interface {
	gobble()
	fly()
}

type realTurkey struct{}

func newTurkey() Turkey {
	return &realTurkey{}
}

func (rT realTurkey) gobble() {
	fmt.Println("Turkey Gobble")
}

func (rT realTurkey) fly() {
	fmt.Println("Turkey Fly")
}
