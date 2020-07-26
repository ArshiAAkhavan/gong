package main

import (
	"gong/shell"
	"log"
)

func main() {
	shell := shell.New()
	shell.Start()
}

func panicAttack(message string, err error) {
	if err != nil {
		log.Fatalf("%s \n\tcaused by error:%v\n", message, err)
	}
	return
}
