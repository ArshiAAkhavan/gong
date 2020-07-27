package shell

import (
	"bufio"
	"fmt"
	"gong/command"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Shell struct {
	commandList []command.Command
}

func New() *Shell {
	return &Shell{}
}

func (sh *Shell) Start() {
	stdinReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("gong> ")
		input, _ := stdinReader.ReadString('\n')
		commandline := strings.Trim(string(input), " \n")
		args := strings.Fields(commandline)

		run(args)
	}

}

func (sh *Shell) AddCommand(c *command.Command) {
	sh.commandList = append(sh.commandList, *c)
}

func (sh *Shell) getCommandByName(n string) *command.Command {
	for _, command := range sh.commandList {
		if command.Name() == n {
			return &command
		}
	}
	return nil
}

func run(args []string) {
	command := exec.Command(args[0], args[1:]...)
	outStream, _ := command.StdoutPipe()
	errStream, _ := command.StderrPipe()

	command.Start()

	output, _ := ioutil.ReadAll(outStream)
	errput, _ := ioutil.ReadAll(errStream)

	command.Wait()
	log.Println(string(errput))
	fmt.Println(string(output))
}
