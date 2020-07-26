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
		// panicAttack("failed to read from input", err)
		commandline := strings.Trim(string(input), " \n")
		args := strings.Fields(commandline)

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

}
