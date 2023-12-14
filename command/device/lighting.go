package device

import "fmt"

type Lighting struct {
	location string
	status   string
}

func NewLighting(location string) *Lighting {
	l := &Lighting{location: location}
	return l
}

func (l *Lighting) On() {
	l.status = "on"
	fmt.Println(l.location + " light is " + l.status)
}

func (l *Lighting) Off() {
	l.status = "off"
	fmt.Println(l.location + " light is " + l.status)
}
