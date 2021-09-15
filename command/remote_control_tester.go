package main

import "fmt"

func main() {
	remoteControl := newRemoteControl()

	livingRoomLight := &light{roomName: "living room"}
	kitchenLight := 	&light{roomName: "kichen light"}
	ceilingFan := &ceilingFan{roomName: "living room"}
	garageDoor := &garage{}
	stereo := &stereo{"living room"}

	livingRoomLightOn := newLightOnCommand(livingRoomLight)
	livingRoomLightOff := newLightOffCommand(livingRoomLight)
	kitchenLightOn := newLightOnCommand(kitchenLight)
	kitchenLightOff := newLightOffCommand(kitchenLight)

	ceilingFanOnCommand := newCeilingFanOnCommand(ceilingFan)
	ceilingFanOffCommand := newCeilingFanOffCommand(ceilingFan)

	garageDoorOpen := newGarageDoorOpenCommand(garageDoor)
	garageDoorClose := newGarageDoorCloseCommand(garageDoor)

	stereoOnWithCD := newStereoOnWithCDCommand(stereo)
	stereoOffWithCD := newStereoOffCommand(stereo)

	remoteControl.setCommand(0, livingRoomLightOn, livingRoomLightOff)
	remoteControl.setCommand(1, kitchenLightOn, kitchenLightOff)
	remoteControl.setCommand(2, ceilingFanOnCommand, ceilingFanOffCommand)
	remoteControl.setCommand(3, garageDoorOpen, garageDoorClose)
	remoteControl.setCommand(4, stereoOnWithCD, stereoOffWithCD)

	fmt.Println(remoteControl)

	//  We step through each slot and push its On and Off button.
	remoteControl.onButtonWasPushed(0)
	remoteControl.offButtonWasPushed(0)
	remoteControl.onButtonWasPushed(1)
	remoteControl.offButtonWasPushed(1)
	remoteControl.onButtonWasPushed(2)
	remoteControl.offButtonWasPushed(2)
	remoteControl.onButtonWasPushed(3)
	remoteControl.offButtonWasPushed(3)
	remoteControl.onButtonWasPushed(4)
	remoteControl.offButtonWasPushed(4)

	// Test macro commands
	partyOn := []iCommand{garageDoorOpen, kitchenLightOn, stereoOnWithCD}
	partyOff := []iCommand{stereoOffWithCD, kitchenLightOff, garageDoorClose}

	partyOnMacro := newMacroCommand(partyOn)
	partyOffMacro := newMacroCommand(partyOff)

	remoteControl.setCommand(5, partyOnMacro, partyOffMacro)
	remoteControl.onButtonWasPushed(5)
	remoteControl.offButtonWasPushed(5)
}
