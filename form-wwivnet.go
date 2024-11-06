package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Application struct to hold form data
type Application struct {
	Alias           string `json:"alias"`
	Email           string `json:"email"`
	Country         string `json:"country"`
	CityState       string `json:"city_state"`
	AreaCode        string `json:"area_code"`
	BBSName         string `json:"bbs_name"`
	BBSURL          string `json:"bbs_url"`
	BBSPort         int    `json:"bbs_port"`
	BBSSoftware     string `json:"bbs_software"`
	BinkPort        int    `json:"bink_port"`
	ApplicationDate string `json:"application_date"`
	Approved        string `json:"approved"`      // "no" by default
	DateApproved    string `json:"date_approved"` // Empty by default
	LastEdited      string `json:"last_edited"`   // Empty by default
}

// Prompt user for input and ensure non-empty responses
func prompt(label string) string {
	var input string
	for {
		fmt.Printf("%s: ", label)
		fmt.Scanln(&input)
		input = strings.TrimSpace(input)
		if input != "" {
			break
		}
		fmt.Println("This field is required. Please enter a value.")
	}
	return input
}

// Prompt user for an integer input
func promptInt(label string) int {
	var input int
	for {
		fmt.Printf("%s: ", label)
		_, err := fmt.Scanf("%d\n", &input)
		if err == nil {
			break
		}
		fmt.Println("Invalid input. Please enter a valid number.")
	}
	return input
}

// Save application data to a JSON file
func saveApplication(data Application, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("could not encode data to JSON: %w", err)
	}

	return nil
}

func app_wwiv() {
	// Collect application data
	fmt.Println("GHOSTnet WWIVnet Application Form")
	fmt.Println("---------------------------------")

	application := Application{
		Alias:           prompt("Alias, Name, or Handle"),
		Email:           prompt("Email Address"),
		Country:         prompt("Country"),
		CityState:       prompt("City and State"),
		AreaCode:        prompt("Telephone Area Code"),
		BBSName:         prompt("BBS Name"),
		BBSURL:          prompt("BBS URL"),
		BBSPort:         promptInt("BBS Port Number"),
		BBSSoftware:     prompt("BBS Software"),
		BinkPort:        promptInt("Bink Port"),
		ApplicationDate: time.Now().Format("2006-01-02"),
		Approved:        "no", // Default to "no"
		DateApproved:    "",   // Empty by default
		LastEdited:      "",   // Empty by default
	}

	// Save application data to JSON
	filename := "GHOSTnet-WWIVnet-application.json"
	err := saveApplication(application, filename)
	if err != nil {
		fmt.Printf("Error saving application: %v\n", err)
		return
	}

	fmt.Printf("Application saved successfully to %s\n", filename)
}
