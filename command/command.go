package command

type Command struct {
	raw string
}

func New() *Command {
	return &Command{}
}

func (c *Command) Raw() string {
	return c.raw
}
