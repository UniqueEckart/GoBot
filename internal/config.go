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
<<<<<<< HEAD
	WelcomeChannel string `json:"WelcomeChannel"`
	LogChannel     string `json:"LogChannel"`
	LeaveChannel   string `json:"LeaveChannel"`
=======
	WelcomeChannel string `json:"Welcome_Channel_ID"`
	LeaveChannel   string `json:"Leave_Channel_ID"`
>>>>>>> 97e918001166bca123c923229e307395197f3469
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
