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

	firstRune := consumed[0]
	if firstRune == '"' || firstRune == '\'' {
		quote := firstRune
		endOfQuoteIdx := findEndOfQuoteIdx(consumed, quote)
		if endOfQuoteIdx == -1 {
			return string(consumed), []rune{}
		}

		return string(consumed[1:endOfQuoteIdx]), consumed[endOfQuoteIdx+2:]
	}

	endOfWordIdx := findFirstOf(consumed, isWhitespace)
	if endOfWordIdx == -1 {
		return string(consumed), []rune{}
	}

	return string(consumed[:endOfWordIdx]), consumed[endOfWordIdx+1:]
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

func findFirstOf(input []rune, predicate func(rune) bool) int {
	for i := 0; i < len(input); i++ {
		if predicate(input[i]) {
			return i
		}
	}

	return -1
}

func findEndOfQuoteIdx(input []rune, quote rune) int {
	previousRune := input[0]

	for i := 1; i < len(input); i++ {
		currentRune := input[i]
		if currentRune == quote && previousRune != '\\' {
			return i
		}

		previousRune = currentRune
	}

	return -1
}
