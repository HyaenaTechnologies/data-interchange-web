package utility

import (
	"fmt"
)

// Print Help Command Output
func PrintHelp() {
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
}
