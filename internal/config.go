package internal

import (
	"encoding/json"
	"os"
)

var Bot_Config *Config

type Config struct {
	Token          string `json:"token"`
	Prefix         string `json:"prefix"`
	Debug          string `json:"debug"`
	MYSQL_HOST     string `json:"MYSQL_HOST"`
	MYSQL_USER     string `json:"MYSQL_USER"`
	MYSQL_PASSWORD string `json:"MYSQL_PASSWORD"`
	MYSQL_DATABASE string `json:"MYSQL_DATABASE"`
	WelcomeChannel string `json:"WelcomeChannel"`
	LogChannel     string `json:"LogChannel"`
	LeaveChannel   string `json:"LeaveChannel"`
}

func ParseConfigFromJSONFile(fileName string) (c *Config, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		Log("Error creating config file", 1)
		return
	}
	c = new(Config)
	err = json.NewDecoder(f).Decode(c)
	return
}

func ValidateConfig() {
	if Bot_Config.Token == "" {
		Log("Please Provide a Discord Token", 1)
	}
	if Bot_Config.WelcomeChannel == "" {
		Log("Please Provide a Channel ID for the Welcome Channel", 1)
	}
	if Bot_Config.LeaveChannel == "" {
		Log("Please Provide a Channel ID for the Leave Channel", 1)
	}
	if Bot_Config.LogChannel == "" {
		Log("Please Provide a Channel ID for the Bot Log Channel", 1)
	}
}
