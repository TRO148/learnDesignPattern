package command

import "learn/command/device"

type LightOn struct {
	Lighting *device.Lighting
}

func (loc *LightOn) Execute() {
	loc.Lighting.On()
}

func (loc *LightOn) Undo() {
	loc.Lighting.Off()
}
