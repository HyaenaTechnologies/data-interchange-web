package command

import (
	"bufio"
	"data-interchange-server/source/network"
	"data-interchange-server/source/utility"
	"fmt"
	"os"
)

var userInput *bufio.Scanner = bufio.NewScanner(os.Stdin)

var userInputText string = userInput.Text()

var cleanUserInput []string = TokenizeArgumentInput(userInputText)

var inputArguments string = cleanUserInput[0]

// Command Line User Input
func UserInput() {
	for {
		fmt.Println("Enter Command: ")

		userInput.Scan()

		if len(cleanUserInput) == 0 {
			continue
		}

		switch inputArguments {
		case "exit":
			os.Exit(0)
		case "help":
			utility.PrintHelp()
		case "serve":
			network.ServeWebApplication()
		case "version":
			utility.PrintVersionNumber()
		default:
			fmt.Println("Invalid Command: ", inputArguments)
		}
	}
}
