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

var builtinsByName = map[string]func(Command){
	"exit": exitCommand,
	"echo": echoCommand,
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command := promptCommand()
		commandFn, found := builtinsByName[command.Name]
		if !found {
			fmt.Println(command.Name + ": command not found")
			continue
		}

		commandFn(command)
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
