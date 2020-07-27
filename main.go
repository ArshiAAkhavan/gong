package main

import (
	"bufio"
	"fmt"
	"gong/command"
	"gong/shell"
	"os"
	"strings"
	"time"
)

func main() {
	shell := shell.New()
	generateCommands(shell)
	shell.Start()
}

func generateCommands(shell *shell.Shell) {

	shell.AddCommand(command.New("ls", func(args []string) {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		fmt.Println(strings.Fields(input)[1])
		fmt.Println("aliali")
	}))

	shell.AddCommand(command.New("wait", func(args []string) {
		time.Sleep(5 * time.Second)
		fmt.Println("done waiting")
		time.Sleep(5 * time.Second)
		fmt.Println("we are good to go!")
	}))
}
