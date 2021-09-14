//Invoker
package main

type simpleRemoteControl struct {
	slot iCommand
}
/**
 * We have a method for setting
 * the command the slot is going
 * to control. This could be called
 * multiple times if the client of
 * this code wanted to change the
 * behavior of the remote button.
 */
func (r *simpleRemoteControl) setCommand(command iCommand)  {
	r.slot = command
}

func (r *simpleRemoteControl) buttonWasPressed()  {
	r.slot.execute()
}
