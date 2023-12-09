package main

import "fmt"

type CurrentConditionDisplay struct {
	temperature float32
	humidity    float32
}

func (ccd *CurrentConditionDisplay) update(temp float32, humidity float32, pressure float32) {
	ccd.temperature = temp
	ccd.humidity = humidity
	ccd.display()
}

func (ccd *CurrentConditionDisplay) display() {
	fmt.Printf("Current condition: %0.2fF degrees and %0.2f%% humidity\n", ccd.temperature, ccd.humidity)
}
