package command

import (
	"strings"
)

// Tokenize Command Line Arguments from User Input
func TokenizeArgumentInput(input string) []string {
	var toLowerCase string = strings.ToLower(input)

	var inputCommands []string = strings.Fields(toLowerCase)

	return inputCommands
}
