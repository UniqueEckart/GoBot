package internal

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token          string `json:"token"`
	Prefix         string `json:"prefix"`
	MYSQL_HOST     string `json:"MYSQL_HOST"`
	MYSQL_USER     string `json:"MYSQL_USER"`
	MYSQL_PASSWORD string `json:"MYSQL_PASSWORD"`
	MYSQL_DATABASE string `json:"MYSQL_DATABASE"`
	WelcomeChannel string `json:"welcome_channel"`
}

func ParseConfigFromJSONFile(fileName string) (c *Config, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	c = new(Config)
	err = json.NewDecoder(f).Decode(c)
	return
}
