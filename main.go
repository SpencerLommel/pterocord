package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	BotToken string `json:"bot_token"`
	APIURL string `json:"api_url"`
	APIKey string `json:"api_key"`
}

func loadConfig() (*Config, error) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
}
