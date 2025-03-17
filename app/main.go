package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

var builtinsByName = map[string]func(Command){}

func main() {
	builtinsByName = map[string]func(Command){
		"exit": exitCommand,
		"echo": echoCommand,
		"type": typeCommand,
	}

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

func typeCommand(command Command) {
	var commandToCheck string
	if len(command.Args) > 0 {
		commandToCheck = command.Args[0]
	} else {
		commandToCheck = ""
	}

	if _, found := builtinsByName[commandToCheck]; found {
		fmt.Printf("%s is a shell builtin\n", commandToCheck)
		return
	}

	resolved, err := ResolveExecutable(commandToCheck)
	if err != nil && !errors.Is(err, ErrNotFound) {
		fmt.Printf("error resolving executable: %v\n", err)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("%s: not found\n", commandToCheck)
		return
	}

	fmt.Printf("%s is %s\n", commandToCheck, resolved.Path)
}
