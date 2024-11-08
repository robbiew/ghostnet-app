package main

import (
	"encoding/json"
	"fmt"
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

// Save application data to JSON file
func saveApplication(data Application, filename string) error {
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		return fmt.Errorf("could not create data directory: %w", err)
	}

	filePath := fmt.Sprintf("data/%s", filename)
	var applications []Application

	if _, err := os.Stat(filePath); err == nil {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("could not open file: %w", err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&applications); err != nil {
			return fmt.Errorf("could not decode existing data: %w", err)
		}
	}

	applications = append(applications, data)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(applications)
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
	err := saveApplication(application, filename)
	if err != nil {
		fmt.Printf("\r\nError saving application: %v\r\n", err)
	} else {
		fmt.Print("\r\nApplication saved successfully\r\n")
	}
}
