package command

import "learn/command/device"

type LightOff struct {
	Lighting *device.Lighting
}

func (loc *LightOff) Execute() {
	loc.Lighting.On()
}

func (loc *LightOff) Undo() {
	loc.Lighting.Off()
}
