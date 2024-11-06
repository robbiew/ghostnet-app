package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Main menu loop
func mainMenu(dropFilePath string) {
	for {
		fmt.Println(ansi.EraseScreen)
		fmt.Print("Main Menu\r\n")
		fmt.Print("-------------------------\r\n")
		fmt.Print("1. Display Drop File Data\r\n")
		fmt.Print("Q. Quit\r\n")
		fmt.Print("\r\nSelect an option: ")

		// Get keyboard input
		input, err := GetKeyboardInput()
		if err != nil {
			fmt.Printf("\r\nError reading input: %v\r\n", err)
			continue
		}

		// Handle menu options
		switch strings.ToUpper(input) {
		case "1":
			data, err := GetDropFileData(dropFilePath)
			if err != nil {
				fmt.Printf("\r\nError reading drop file data: %v\r\n", err)
			} else {
				fmt.Println(ansi.EraseScreen)
				fmt.Print("Drop File Data:\r\n")
				fmt.Print("-------------------------\r\n")
				fmt.Printf(" - Alias: %s\r\n", data.Alias)
				fmt.Printf(" - Time Left: %d minutes\r\n", data.TimeLeft)
				fmt.Printf(" - Emulation: %d\r\n", data.Emulation)
				fmt.Printf(" - Node Number: %d\r\n\r\n", data.NodeNum)
				Pause()
			}
		case "Q":
			fmt.Println("\r\nExiting program.")
			return
		default:
			fmt.Println("\r\nInvalid option. Please select 1 or Q.")
		}
	}
}

func main() {
	// Define the command-line flag
	dropFilePath := flag.String("path", "", "Path to the directory containing door32.sys")
	flag.Parse()

	// Check if the path flag is provided
	if *dropFilePath == "" {
		fmt.Println("Please provide the path to the drop file directory using the -path flag.")
		os.Exit(1)
	}

	fmt.Println(ansi.EraseScreen)

	// Start main menu loop
	mainMenu(*dropFilePath)

}
