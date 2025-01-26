package command

import (
	"bufio"
	"data-interchange-server/source/network"
	"data-interchange-server/source/utility"
	"fmt"
	"os"
	"strings"
)

// Tokenize Command Line Arguments from User Input
func TokenizeArgumentInput(input string) []string {
	if len(input) == 0 {
		fmt.Println("Input Required")
		utility.ErrorExit()
	}

	var toLowerCase string = strings.ToLower(input)

	var inputCommands []string = strings.Fields(toLowerCase)

	return inputCommands
}

func temp() {
	var userInput *bufio.Scanner = bufio.NewScanner(os.Stdin)

	var userInputText string = userInput.Text()

	var tokenizedUserInput []string = TokenizeArgumentInput(userInputText)

	var inputArguments string = tokenizedUserInput[0]

	for {
		fmt.Println("Enter Command: ")

		userInput.Scan()

		if len(tokenizedUserInput) == 0 {
			continue
		}

		var availableCommands map[string]InputCommand = CommandMap()

		var inputCommand InputCommand = availableCommands[inputArguments]

		inputCommand.event()

		switch inputArguments {
		case "exit":
			utility.SuccesfulExit()
		case "help":
			utility.PrintHelp()
		case "serve":
			network.ServeWebApplication()
		case "version":
			utility.PrintVersionNumber()
		default:
			fmt.Println("Invalid Command: ", inputArguments)
			utility.ErrorExit()
		}
	}
}
