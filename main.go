package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {

	stdinReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$> ")
		input, _ := stdinReader.ReadString('\n')
		commandline := strings.Trim(string(input), " \n")

		command := exec.Command(commandline)
		outStream, _ := command.StdoutPipe()
		command.Start()
		out, _ := ioutil.ReadAll(outStream)
		command.Wait()
		fmt.Println(string(out))
	}
}
