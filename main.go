package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/SpencerLommel/pterocord/util"
	"github.com/bwmarrin/discordgo"
)

type Config struct {
	BotToken string `json:"bot_token"`
	GuildID  string `json:"guild_id"`
}

var (
	s      *discordgo.Session
	config *util.Config
)

func whitelistCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name != "whitelist" {
		return
	}

	username := i.ApplicationCommandData().Options[0].StringValue()
	discordUser := i.Member.User.Username

	fmt.Printf("@%s tried to whitelist %s\n", discordUser, username)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Attempting to whitelist %s...", username),
		},
	})
}

func registerCommands(s *discordgo.Session, guildID string) ([]*discordgo.ApplicationCommand, error) {
	cmds := []*discordgo.ApplicationCommand{
		{
			Name:        "whitelist",
			Description: "Whitelist a Minecraft user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Minecraft username to whitelist",
					Required:    true,
				},
			},
		},
	}

	var registeredCommands []*discordgo.ApplicationCommand
	for _, cmd := range cmds {
		createdCmd, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, cmd)
		if err != nil {
			return nil, fmt.Errorf("failed to create command %s: %w", cmd.Name, err)
		}
		registeredCommands = append(registeredCommands, createdCmd)
	}

	return registeredCommands, nil
}

func cleanupCommands(s *discordgo.Session, guildID string, commands []*discordgo.ApplicationCommand) {
	for _, cmd := range commands {
		err := s.ApplicationCommandDelete(s.State.User.ID, guildID, cmd.ID)
		if err != nil {
			fmt.Printf("Failed to delete command %s: %v\n", cmd.Name, err)
		} else {
			fmt.Printf("Deleted command %s\n", cmd.Name)
		}
	}
}

func main() {
	var err error
	config, err = util.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	s, err = discordgo.New("Bot " + config.BotToken)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	s.AddHandler(whitelistCommandHandler)
	s.Identify.Intents = discordgo.IntentsGuilds

	if err = s.Open(); err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	defer s.Close()

	commands, err := registerCommands(s, config.GuildID)
	if err != nil {
		fmt.Println("Error registering commands:", err)
		return
	}

	fmt.Println("Bot is running. Press CTRL+C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	cleanupCommands(s, config.GuildID, commands)
}
