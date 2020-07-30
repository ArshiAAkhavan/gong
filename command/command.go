package command

type Executable interface {
	Name() string
	Help() string
	SetArgs([]string)
	Execute()
}

type Command struct {
	name     string
	help     string
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

func (c *Command) Help() string {
	return c.help
}

func (c *Command) Execute() {
	c.function(c.args)
}
