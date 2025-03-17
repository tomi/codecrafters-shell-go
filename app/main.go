package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command := promptCommand()

		fmt.Println(command + ": command not found")
	}
}

func promptCommand() string {
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading command: %v\n", err)
		os.Exit(1)
	}

	return command[:len(command)-1]
}
