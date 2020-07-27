package main

import (
	"bufio"
	"fmt"
	"gong/command"
	"gong/shell"
	"os"
	"strings"
)

func main() {
	shell := shell.New()
	shell.AddCommand(command.New("ls", func(args []string) {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		fmt.Println(strings.Fields(input)[1])
		fmt.Println("aliali")
	}))
	shell.Start()
}
