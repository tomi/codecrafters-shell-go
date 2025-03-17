package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

var builtinsByName = map[string]func(Command){}
var navigator *Navigator

func main() {
	navigator = MakeNavigator()
	builtinsByName = map[string]func(Command){
		"exit": exitCommand,
		"echo": echoCommand,
		"type": typeCommand,
		"pwd":  pwdCommand,
		"cd":   cdCommand,
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		command := promptCommand()
		commandFn, found := builtinsByName[command.Name]
		if found {
			commandFn(command)
			continue
		}

		executable, err := ResolveExecutable(command.Name)
		if err != nil && !errors.Is(err, ErrNotFound) {
			fmt.Printf("error resolving executable: %v\n", err)
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(command.Name + ": command not found")
			continue
		}

		runExecutable(executable, command)
	}
}

func promptCommand() Command {
	commandStr, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading command: %v\n", err)
		os.Exit(1)
	}

	command, err := ParseInput(commandStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing command: %v\n", err)
		os.Exit(1)
	}

	return command
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

func runExecutable(exe ResolvedExecutable, command Command) {
	cmd := exec.Command(exe.Name, command.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func pwdCommand(_ Command) {
	navigator.PrintWorkingDirectory()
}

func cdCommand(command Command) {
	if len(command.Args) == 0 {
		fmt.Printf("cd: Expected argument\n")
		return
	}

	dir := command.Args[0]
	err := navigator.ChangeDirectory(dir)
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("cd: %s: No such file or directory\n", dir)
		return
	}

	if errors.Is(err, ErrNotADirectory) {
		fmt.Printf("cd: %s: Not a directory\n", dir)
		return
	}

	if err != nil {
		fmt.Printf("cd: %s: Unexpected error: %v\n", dir, err)
		return
	}
}
