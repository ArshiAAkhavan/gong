package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	stdinReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$> ")
		input, err := stdinReader.ReadString('\n')
		panicAttack("failed to read from input", err)
		commandline := strings.Trim(string(input), " \n")
		args := strings.Fields(commandline)
		command := exec.Command(args[0], args[1:]...)
		outStream, err := command.StdoutPipe()
		errStream, err := command.StderrPipe()
		command.Start()
		output, err := ioutil.ReadAll(outStream)
		errput, _ := ioutil.ReadAll(errStream)

		command.Wait()
		log.Println(string(errput))
		fmt.Println(string(output))
	}
}

func panicAttack(message string, err error) {
	if err != nil {
		log.Fatalf("%s \n\tcaused by error:%v\n", message, err)
	}
	return
}
