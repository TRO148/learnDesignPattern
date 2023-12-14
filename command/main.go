package main

import (
	"fmt"
	"learn/command/command"
	"learn/command/device"
)

func main() {
	livingRoomLight := device.NewLighting("livingRoom")
	kitchenLight := device.NewLighting("kitchen")
	ceilingFan := device.NewFan("ceiling")

	livingRoomLightOn := &command.LightOn{Lighting: livingRoomLight}
	livingRoomLightOff := &command.LightOff{Lighting: livingRoomLight}
	kitchenLightOn := &command.LightOff{Lighting: kitchenLight}
	kitchenLightOff := &command.LightOff{Lighting: kitchenLight}
	ceilingFanHigh := &command.FanHigh{L: ceilingFan}
	ceilingFanLow := &command.FanLow{L: ceilingFan}
	ceilingFanOff := &command.FanOff{L: ceilingFan}

	livingRoomLightOnWithHigh := command.NewGroup(livingRoomLightOn, ceilingFanHigh)
	livingRoomLightOffWithLow := command.NewGroup(livingRoomLightOff, ceilingFanLow)
	livingRoomLightOnWithHighAndOffWithLow := command.NewGroup(livingRoomLightOn, ceilingFanHigh, livingRoomLightOff, ceilingFanLow)

	rc := NewRemoteControl()
	rc.setCommand(0, livingRoomLightOff, livingRoomLightOn, &command.No{})
	rc.setCommand(1, kitchenLightOff, kitchenLightOn, &command.No{})
	rc.setCommand(2, ceilingFanOff, ceilingFanHigh, ceilingFanLow)
	rc.setCommand(3, livingRoomLightOnWithHigh, livingRoomLightOffWithLow, livingRoomLightOnWithHighAndOffWithLow)

	rc.onCommandWasPush(0)
	rc.offCommandWasPush(0)
	rc.onCommandWasPush(1)

	rc.onCommandWasPush(2)
	rc.offCommandWasPush(2)

	rc.onCommandWasPush(3)
	rc.offCommandWasPush(3)
	rc.Undo()

	fmt.Println(rc.ShowCommand())
}
