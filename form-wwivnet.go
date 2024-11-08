package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

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
	Approved        string `json:"approved"`      // Default to "no"
	DateApproved    string `json:"date_approved"` // Empty by default
	LastEdited      string `json:"last_edited"`   // Empty by default
}

// SaveJSON appends a new application to the JSON file in the "data" directory
func SaveJSON(data interface{}, filename string) error {
	// Ensure the "data" directory exists
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		return fmt.Errorf("\r\ncould not create data directory: %w", err)
	}

	// Set the file path to save within the "data" directory
	filePath := fmt.Sprintf("data/%s", filename)

	var applications []Application

	// Read existing applications if the file exists
	if _, err := os.Stat(filePath); err == nil {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("\r\ncould not open file: %w", err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&applications); err != nil && err != io.EOF {
			return fmt.Errorf("\r\ncould not decode existing data: %w", err)
		}
	}

	// Append the new application (assuming data is of type Application)
	if app, ok := data.(Application); ok {
		applications = append(applications, app)
	} else {
		return fmt.Errorf("\r\ndata is not of type Application")
	}

	// Write the updated applications list to the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("\r\ncould not create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(applications); err != nil {
		return fmt.Errorf("\r\ncould not encode data to JSON: %w", err)
	}

	return nil
}

func app_wwiv() {
	fmt.Print("GHOSTnet WWIVnet Application Form\r\n")
	fmt.Print("---------------------------------\r\n")

	application := Application{
		Alias:           Prompt("\r\nAlias, Name, or Handle"),
		Email:           Prompt("\r\nEmail Address"),
		Country:         Prompt("\r\nCountry"),
		CityState:       Prompt("\r\nCity and State"),
		AreaCode:        Prompt("\r\nTelephone Area Code"),
		BBSName:         Prompt("\r\nBBS Name"),
		BBSURL:          Prompt("\r\nBBS URL"),
		BBSPort:         PromptInt("\r\nBBS Port Number"),
		BBSSoftware:     Prompt("\r\nBBS Software"),
		BinkPort:        PromptInt("\r\nBink Port"),
		ApplicationDate: time.Now().Format("2006-01-02"),
		Approved:        "no",
		DateApproved:    "",
		LastEdited:      "",
	}

	filename := "GHOSTnet-WWIVnet-application.json"
	err := SaveJSON(application, filename)
	if err != nil {
		fmt.Printf("\r\nError saving application: %v\r\n", err)
	} else {
		fmt.Print("\r\nApplication saved successfully\r\n")
	}
}
