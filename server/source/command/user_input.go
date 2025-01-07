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
			fmt.Println("Hyaena Technologies: Data Interchange Server")
			fmt.Println("")
			fmt.Println("")
			fmt.Println("Commands:				Description:")
			fmt.Println("")
			fmt.Println("exit 					Exit Server")
			fmt.Println("help					Print List of Commands and Flags")
			fmt.Println("serve 					Serve Web Applcation")
			fmt.Println("version				Print Version Number")
			fmt.Println("")
			fmt.Println("")
			fmt.Println("Flags:				    Description:")
			fmt.Println("")
			fmt.Println("--help, --h, -h		Print List of Commands and Flags")
			fmt.Println("--version, --v, -v		Print Version Number")
		case "serve":
			network.ServeWebApplication()
		case "version":
			utility.PrintVersionNumber()
		default:
			fmt.Println("Invalid Command: ", inputArguments)
		}
	}
}
