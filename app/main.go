package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command := promptCommand()
		switch command.Name {
		case "exit":
			exitCommand(command)

		case "echo":
			echoCommand(command)

		default:
			fmt.Println(command.Name + ": command not found")
		}
	}
}

func promptCommand() Command {
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading command: %v\n", err)
		os.Exit(1)
	}

	parts := strings.Split(command[:len(command)-1], " ")

	return Command{
		Name: parts[0],
		Args: parts[1:],
	}
}

func exitCommand(_ Command) {
	os.Exit(0)
}

func echoCommand(command Command) {
	fmt.Println(strings.Join(command.Args, " "))
}
