package command

import "learn/command/device"

type FanOff struct {
	L            *device.Fan
	beforeStatus string
}

func (loc *FanOff) Execute() {
	loc.beforeStatus = loc.L.GetStatus()
	loc.L.Off()
}

func (loc *FanOff) Undo() {
	// Fan不应提供直接设置的接口, 要不就算设置状态了, 所以这里采用这种繁重的方式
	switch loc.beforeStatus {
	case "high":
		loc.L.High()
	case "low":
		loc.L.Low()
	default:
		loc.L.Off()
	}
}
