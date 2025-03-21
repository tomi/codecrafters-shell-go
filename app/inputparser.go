package main

import (
	"fmt"
)

type ParserQuoteState int

const (
	ParserQuoteStateNone ParserQuoteState = iota
	ParserQuoteStateInSingleQuote
	ParserQuoteStateInDoubleQuote
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

	quoteState := ParserQuoteStateNone
	isEscaped := false
	word := []rune{}
	idx := 0

	for idx < len(consumed) {
		currentRune := consumed[idx]

		if quoteState == ParserQuoteStateInSingleQuote {
			if currentRune == '\'' {
				quoteState = ParserQuoteStateNone
			} else {
				word = append(word, currentRune)
			}
		} else if quoteState == ParserQuoteStateInDoubleQuote {
			if isEscaped {
				// Only ", \ and $ have special meaning
				if currentRune == '"' || currentRune == '\\' || currentRune == '$' {
					word = append(word, currentRune)
				} else {
					word = append(word, '\\', currentRune)
				}
				isEscaped = false
			} else if currentRune == '"' {
				quoteState = ParserQuoteStateNone
			} else if currentRune == '\\' {
				isEscaped = true
			} else {
				word = append(word, currentRune)
			}
		} else if isEscaped {
			isEscaped = false
			word = append(word, currentRune)
		} else if currentRune == '\\' {
			isEscaped = true
		} else if currentRune == '"' {
			quoteState = ParserQuoteStateInDoubleQuote
		} else if currentRune == '\'' {
			quoteState = ParserQuoteStateInSingleQuote
		} else if isWhitespace(currentRune) && !isEscaped {
			break
		} else {
			isEscaped = false
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
