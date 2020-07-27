package command

type Command struct {
	name     string
	args     []string
	function func([]string)
}

func New(name string, function func([]string)) *Command {
	return &Command{
		name:     name,
		function: function,
	}
}

func (c *Command) SetArgs(args []string) {
	c.args = args
}

func (c *Command) Name() string {
	return c.name
}

func (c *Command) Run() {
	c.function(c.args)
}
