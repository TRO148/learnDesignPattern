package device

import "fmt"

type Fan struct {
	location string
	level    string
}

func NewFan(location string) *Fan {
	f := &Fan{level: "high"}
	return f
}

func (f *Fan) GetStatus() string {
	return f.level
}

func (f *Fan) High() {
	f.level = "high"
	fmt.Println(f.location + "Fan is " + f.level)

}

func (f *Fan) Low() {
	f.level = "low"
	fmt.Println(f.location + "Fan is " + f.level)

}

func (f *Fan) Off() {
	f.level = "off"
	fmt.Println(f.location + " Fan is " + f.level)
}
