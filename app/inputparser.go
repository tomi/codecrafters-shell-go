package main

import (
	"fmt"
)

func ParseInput(input string) (Command, error) {
	if len(input) == 0 {
		return Command{}, fmt.Errorf("empty input")
	}

	inputAsRunes := []rune(input)
	var commandName string
	var args []string

	commandName, restOfInput := consumeWord(inputAsRunes)
	if commandName == "" {
		return Command{}, fmt.Errorf("empty command name")
	}

	if len(restOfInput) == 0 {
		return Command{
			Name: commandName,
			Args: []string{},
		}, nil
	}

	for {
		arg, restOfArgs := consumeWord(restOfInput)
		if arg == "" {
			break
		}

		args = append(args, arg)
		restOfInput = restOfArgs
	}

	return Command{
		Name: commandName,
		Args: args,
	}, nil
}

func consumeWord(input []rune) (string, []rune) {
	consumed := consumeWhitespace(input)
	if len(consumed) == 0 {
		return "", consumed
	}

	quote := rune(0)
	word := []rune{}
	idx := 0

	for idx < len(consumed) {
		currentRune := consumed[idx]
		if currentRune == quote {
			quote = rune(0)
		} else if currentRune == '"' || currentRune == '\'' {
			quote = currentRune
		} else if isWhitespace(currentRune) && quote == rune(0) {
			break
		} else {
			word = append(word, currentRune)
		}

		idx++
	}

	return string(word), consumed[idx:]
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func consumeWhitespace(input []rune) []rune {
	idxOfNonWhitespace := 0
	for idxOfNonWhitespace < len(input) {
		if !isWhitespace(input[idxOfNonWhitespace]) {
			break
		}

		idxOfNonWhitespace++
	}

	return input[idxOfNonWhitespace:]
}
