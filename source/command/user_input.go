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

var tokenizedUserInput []string = TokenizeArgumentInput(userInputText)

// Command Line User Input
func UserInput() {
	for {
		fmt.Println("Enter Command: ")

		userInput.Scan()

		if len(tokenizedUserInput) == 0 {
			continue
		}

		switch tokenizedUserInput[0] {
		case "exit":
			os.Exit(0)
		case "help":
			utility.PrintHelp()
		case "serve":
			network.ServeWebApplication()
		case "version":
			utility.PrintVersionNumber()
		default:
			fmt.Println("Invalid Command: ", tokenizedUserInput)
		}
	}
}
