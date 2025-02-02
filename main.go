package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	BotToken string `json:"bot_token"`
	GuildID  string `json:"guild_id"`
	APIURL   string `json:"api_url"`
	APIKey   string `json:"api_key"`
}

var s *discordgo.Session

func loadConfig() (*Config, error) {
	data, err := os.ReadFile("config.json")
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
	s, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

}
