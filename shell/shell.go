package shell

import (
	"bufio"
	"fmt"
	"gong/command"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

type Shell struct {
	execs     []command.Executable
	os_sigs   chan os.Signal
	done_sigs chan int
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
			input, err := stdinReader.ReadString('\n')
			if err == io.EOF {
				log.Fatalf("ali ali")
			}
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

func (sh *Shell) AddExecutable(exec command.Executable) {
	sh.execs = append(sh.execs, exec)
}

func (sh *Shell) getExecutableByName(n string) *command.Executable {
	for _, exec := range sh.execs {
		if exec.Name() == n {
			return &exec
		}
	}
	return nil
}

//todo Ctrl+C doesnt work properly , check  github.com/matryer/runner
func (sh *Shell) run(args []string) {
	if len(args) == 0 {
		fmt.Println("bye!")
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("command entered an error state!")
		}
	}()
	x := *sh.getExecutableByName(args[0])

	fmt.Println(x.Name())
	if x == nil {
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
	x.SetArgs(args)
	x.Execute()

}
