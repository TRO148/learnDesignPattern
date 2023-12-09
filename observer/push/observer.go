package main

type Observer interface {
	update(temp float32, humidity float32, pressure float32)
}
