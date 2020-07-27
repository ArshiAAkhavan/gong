package shell

import (
	"bufio"
	"fmt"
	"gong/command"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

type Shell struct {
	commandList []command.Command
	os_sigs     chan os.Signal
	done_sigs   chan int
}

func New() *Shell {
	sh := Shell{}
	sh.done_sigs = make(chan int, 1)
	sh.os_sigs = make(chan os.Signal, 1)
	signal.Notify(sh.os_sigs, syscall.SIGINT, syscall.SIGTERM)

	return &sh
}

func (sh *Shell) Start() {
	stdinReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("gong> ")

		go func() {
			input, _ := stdinReader.ReadString('\n')
			commandline := strings.Trim(string(input), " \n")
			args := strings.Fields(commandline)

			sh.run(args)
			sh.done_sigs <- 0
		}()
		select {
		case intrupt_signal := <-sh.os_sigs:
			log.Println(intrupt_signal.String)
		case exit_code := <-sh.done_sigs:
			_ = exit_code
		}
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

func (sh *Shell) run(args []string) {
	c := sh.getCommandByName(args[0])
	if c == nil {
		command := exec.Command(args[0], args[1:]...)
		outStream, _ := command.StdoutPipe()
		errStream, _ := command.StderrPipe()

		command.Start()

		output, _ := ioutil.ReadAll(outStream)
		errput, _ := ioutil.ReadAll(errStream)

		command.Wait()
		log.Println(string(errput))
		fmt.Println(string(output))
		return
	}

	c.SetArgs(args)
	c.Run()

}
