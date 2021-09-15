package main

func mains() {
	/**
	 * The remote is our Invoker;
	 * it will be passed a
	 * command object that can
	 * be used to make requests.
	 */
	remote := &simpleRemoteControl{}

	/**
	 * Now we create a Light
	 * object, this will be the
	 * Receiver of the request
	 */
	light  := &light{roomName: "My Room"}

	/**
	 * Create a command and
	 * pass the Receiver to it.
	 */
	lightOnCommand := newLightOnCommand(light)

	/**
	 * Pass the command to the Invoker
	 */
	remote.setCommand(lightOnCommand)
	/**
	 * We simulate the button being pressed.
	 */
	remote.buttonWasPressed()

	garage := &garage{}
	garageDoorOpenCommand := newGarageDoorOpenCommand(garage)

	remote.setCommand(garageDoorOpenCommand)
	remote.buttonWasPressed()
}
