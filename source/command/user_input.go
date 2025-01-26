package command

import (
	"data-interchange-server/source/network"
	"data-interchange-server/source/utility"
	"fmt"
	"os"
)

// Command Line User Input
func UserInput() {
	var inputArguments []string = os.Args[1:]

	var tokenizedInputArguments []string = TokenizeArgumentInput(inputArguments[0])

	if len(inputArguments) == 0 {
		fmt.Println("Input Required")
		utility.ErrorExit()
	} else {
		switch tokenizedInputArguments[0] {
		case "exit":
			utility.SuccesfulExit()
		case "--exit":
			utility.SuccesfulExit()
		case "-exit":
			utility.SuccesfulExit()
		case "--e":
			utility.SuccesfulExit()
		case "-e":
			utility.SuccesfulExit()
		case "help":
			utility.PrintHelp()
		case "--help":
			utility.PrintHelp()
		case "-help":
			utility.PrintHelp()
		case "--h":
			utility.PrintHelp()
		case "-h":
			utility.PrintHelp()
		case "serve":
			network.ServeWebApplication()
		case "version":
			utility.PrintVersionNumber()
		case "--version":
			utility.PrintVersionNumber()
		case "-version":
			utility.PrintVersionNumber()
		case "--v":
			utility.PrintVersionNumber()
		case "-v":
			utility.PrintVersionNumber()
		default:
			fmt.Println("Invalid Command: ", tokenizedInputArguments)
			utility.ErrorExit()
		}
	}
}
