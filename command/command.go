package main

type iCommand interface {
	execute()
}

type lightOnCommand struct {
	*light
}

func newLightOnCommand(l *light) *lightOnCommand {
	return &lightOnCommand{
		light: l,
	}
}

func (l *lightOnCommand) execute()  {
	l.light.on()
}

type lightOffCommand struct {
	*light
}

func newLightOffCommand(l *light) *lightOffCommand {
	return &lightOffCommand{
		l,
	}
}

func (l *lightOffCommand) execute()  {
	l.light.off()
}

type garageDoorOpenCommand struct {
	*garage
}

func newGarageDoorOpenCommand(g *garage) *garageDoorOpenCommand{
	return &garageDoorOpenCommand{
		g,
	}
}

func (g *garageDoorOpenCommand) execute()  {
	g.garage.up()
}

type garageDoorCloseCommand struct {
	*garage
}

func newGarageDoorCloseCommand(g *garage) *garageDoorCloseCommand{
	return &garageDoorCloseCommand{
		g,
	}
}

func (g *garageDoorCloseCommand) execute()  {
	g.garage.down()
}

type stereoOnWithCDCommand struct {
	*stereo
}

func newStereoOnWithCDCommand(c *stereo) *stereoOnWithCDCommand{
	return &stereoOnWithCDCommand{
		c,
	}
}

func (s *stereoOnWithCDCommand) execute()  {
	s.stereo.on()
	s.stereo.setCD()
	s.stereo.setVolume(11)
}

type stereoOffCommand struct {
	stereo *stereo
}

func newStereoOffCommand(s *stereo) *stereoOffCommand {
	return &stereoOffCommand{
		stereo: s,
	}
}

func (s *stereoOffCommand) execute() {
	s.stereo.off()
}

type ceilingFanOnCommand struct {
	*ceilingFan
}

func (c *ceilingFanOnCommand) execute() {
	c.ceilingFan.on()
}

func newCeilingFanOnCommand(c *ceilingFan) *ceilingFanOnCommand {
	return &ceilingFanOnCommand{
		c,
	}
}

type ceilingFanOffCommand struct {
	*ceilingFan
}

func (c *ceilingFanOffCommand) execute() {
	c.ceilingFan.off()
}

func newCeilingFanOffCommand(c *ceilingFan) *ceilingFanOffCommand {
	return &ceilingFanOffCommand{
		c,
	}
}

type macroCommand struct {
	commands []iCommand
}

func newMacroCommand(commands []iCommand) *macroCommand {
	return &macroCommand{
		commands,
	}
}

func (m *macroCommand) execute()  {
	for _, command := range m.commands {
		command.execute()
	}
}
