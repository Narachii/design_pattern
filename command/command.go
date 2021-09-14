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
