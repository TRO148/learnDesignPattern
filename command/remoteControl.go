package main

import (
	"bytes"
	"fmt"
	"learn/command/command"
	"reflect"
)

type RemoteControl struct {
	onCommand    [7]command.Interface
	otherCommand [7]command.Interface
	offCommand   [7]command.Interface

	beforeUndo func()
}

func NewRemoteControl() *RemoteControl {
	nrc := new(RemoteControl)
	for i := 0; i < 7; i++ {
		nrc.onCommand[i] = &command.No{}
		nrc.offCommand[i] = &command.No{}
		nrc.otherCommand[i] = &command.No{}
	}
	return nrc
}

func (rc *RemoteControl) Undo() {
	if rc.beforeUndo != nil {
		fmt.Print("Undo: ")
		rc.beforeUndo()
	}
}

func (rc *RemoteControl) setCommand(slot int, offCommand, onCommand, otherCommand command.Interface) {
	rc.onCommand[slot] = onCommand
	rc.offCommand[slot] = offCommand
	rc.otherCommand[slot] = otherCommand
}

func (rc *RemoteControl) onCommandWasPush(slot int) {
	rc.onCommand[slot].Execute()
	rc.beforeUndo = rc.onCommand[slot].Undo
}

func (rc *RemoteControl) offCommandWasPush(slot int) {
	rc.offCommand[slot].Execute()
	rc.beforeUndo = rc.offCommand[slot].Undo
}

func (rc *RemoteControl) otherCommandWasPush(slot int) {
	rc.otherCommand[slot].Execute()
	rc.beforeUndo = rc.otherCommand[slot].Undo
}

func (rc *RemoteControl) ShowCommand() string {
	var byt bytes.Buffer
	byt.WriteString("\n-------------Remote Control-------------\n")
	byt.WriteString("\n------On Command Show------\n")
	for i := 0; i < 7; i++ {
		onCommandType := reflect.TypeOf(rc.onCommand[i])
		offCommandType := reflect.TypeOf(rc.offCommand[i])
		otherCommandType := reflect.TypeOf(rc.otherCommand[i])
		byt.WriteString(fmt.Sprintf("[slot %d]\t%s\t\t%s\t\t%s", i,
			onCommandType.Elem().Name(), offCommandType.Elem().Name(), otherCommandType.Elem().Name()))
		byt.WriteString("\n")
	}
	return byt.String()
}
