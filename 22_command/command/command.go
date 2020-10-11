package command

type Command struct {
	*Beaker
}

func (c *Command) SetBeaker(beaker *Beaker) {
	c.Beaker = beaker
}

func (c *Command) Execute() {
}