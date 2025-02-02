package util

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	BotToken string `json:"bot_token"`
	GuildID  string `json:"guild_id"`
	APIURL   string `json:"api_url"`
	APIKey   string `json:"api_key"`
}

// GenerateConfig prompts the user for config values and saves them to config.json
func GenerateConfig() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter bot_token: ")
	botToken, _ := reader.ReadString('\n')
	botToken = strings.TrimSpace(botToken)

	fmt.Print("Enter guild_id: ")
	guildID, _ := reader.ReadString('\n')
	guildID = strings.TrimSpace(guildID)

	fmt.Print("Enter api_url: ")
	apiURL, _ := reader.ReadString('\n')
	apiURL = strings.TrimSpace(apiURL)

	fmt.Print("Enter api_key: ")
	apiKey, _ := reader.ReadString('\n')
	apiKey = strings.TrimSpace(apiKey)

	config := Config{
		BotToken: botToken,
		GuildID:  guildID,
		APIURL:   apiURL,
		APIKey:   apiKey,
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error creating config JSON:", err)
		return
	}

	err = os.WriteFile("config.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing config file:", err)
	}
}

// LoadConfig reads and parses config.json, generating it if it doesn't exist
func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("config.json")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Config doesn't exist, generating new one.")
			GenerateConfig()
			return LoadConfig() // Rerun after generating
		}
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
