package command

type Command struct {
	name string
	args []string
}

func New() *Command {
	return &Command{}
}

func (c *Command) Name() string {
	return c.name
}
