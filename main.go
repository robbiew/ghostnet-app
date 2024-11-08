package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Config struct to hold settings from the ini file
type Config struct {
	AdminSecurityLevel int
	WWIVnet            bool
	FTN                bool
}

// Struct for storing drop file data
type DropFileData struct {
	CommType      int
	CommHandle    int
	BaudRate      int
	BBSID         string
	UserRecordPos int
	RealName      string
	Alias         string
	SecurityLevel int
	TimeLeft      int
	Emulation     int
	NodeNum       int
}

// Main menu loop
func mainMenu(dropFilePath string, config *Config) {
	for {
		fmt.Println(ansi.EraseScreen)
		fmt.Print("Main Menu\r\n")
		fmt.Print("-------------------------\r\n")
		fmt.Print("1. [DEBUG] View Drop File Data\r\n")
		data, err := GetDropFileData(dropFilePath)
		if err != nil {
			fmt.Printf("\r\nError reading drop file data: %v\r\n", err)
			continue
		}
		if data.SecurityLevel >= config.AdminSecurityLevel {
			fmt.Print("2. [DEBUG] Config & User Access Check\r\n")
		}
		fmt.Print("3. Apply for GHOSTnet WWIVnet Node\r\n")
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
			// Display drop file data
			displayDropFileData(data)
		case "2":
			// Display configuration and check access
			displayConfig(config, data)
		case "3":
			// Call the WWIVnet application function
			fmt.Println(ansi.EraseScreen)
			app_wwiv()
			Pause()
		case "Q":
			fmt.Println("\r\nExiting program.")
			return
		default:
			fmt.Println("\r\nInvalid option. Please select 1, 2, 3, or Q.")
		}
	}
}

// Helper function to display drop file data
func displayDropFileData(data DropFileData) {
	fmt.Println(ansi.EraseScreen)
	fmt.Print("DOOR32.SYS Data:\r\n")
	fmt.Print("-------------------------\r\n")
	fmt.Printf(" - Comm Type: %d\r\n", data.CommType)
	fmt.Printf(" - Comm Handle: %d\r\n", data.CommHandle)
	fmt.Printf(" - Baud Rate: %d\r\n", data.BaudRate)
	fmt.Printf(" - BBSID: %s\r\n", data.BBSID)
	fmt.Printf(" - User Record Position: %d\r\n", data.UserRecordPos)
	fmt.Printf(" - Real Name: %s\r\n", data.RealName)
	fmt.Printf(" - Alias: %s\r\n", data.Alias)
	fmt.Printf(" - Security Level: %d\r\n", data.SecurityLevel)
	fmt.Printf(" - Time Left: %d minutes\r\n", data.TimeLeft)
	fmt.Printf(" - Emulation: %d\r\n", data.Emulation)
	fmt.Printf(" - Node Number: %d\r\n\r\n", data.NodeNum)
	Pause()
}

// Helper function to display config and access check
func displayConfig(config *Config, data DropFileData) {
	fmt.Println(ansi.EraseScreen)
	fmt.Print("Configuration Settings:\r\n")
	fmt.Print("-------------------------\r\n")
	fmt.Printf(" - Admin Security Level: %d\r\n", config.AdminSecurityLevel)
	fmt.Printf(" - WWIVnet Enabled: %v\r\n", config.WWIVnet)
	fmt.Printf(" - FTN Enabled: %v\r\n", config.FTN)

	if data.SecurityLevel >= config.AdminSecurityLevel {
		fmt.Print(" - " + data.Alias + " has admin access\r\n")
	} else {
		fmt.Print(" - " + data.Alias + " does not have admin access\r\n")
	}
	fmt.Print("\r\n")
	Pause()
}

func main() {
	// Load configuration
	config, err := LoadConfig("config.ini")
	if err != nil {
		fmt.Printf("\r\nError loading configuration: %v\n", err)
		os.Exit(1)
	}

	// Define the command-line flag
	dropFilePath := flag.String("path", "", "Path to the directory containing door32.sys")
	flag.Parse()

	if *dropFilePath == "" {
		fmt.Println("Please provide the path to the drop file directory using the -path flag.")
		os.Exit(1)
	}

	mainMenu(*dropFilePath, config)
}
