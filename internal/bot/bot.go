// Package bot defines the structure and functionality of the AgoraBot.
package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/maaxleq/agora-bot/internal/config"
)

// AgoraBot represents a Discord bot with configuration and session information.
type AgoraBot struct {
	Conf    config.Config
	Session *discordgo.Session
}

// NewAgoraBot creates a new instance of AgoraBot with the provided configuration.
func NewAgoraBot(conf config.Config) (*AgoraBot, error) {
	dg, errBot := discordgo.New("Bot " + conf.DiscordToken)
	if errBot != nil {
		return nil, fmt.Errorf("error creating Discord session: %w", errBot)
	}

	return &AgoraBot{
		Conf:    conf,
		Session: dg,
	}, nil
}

// Run starts the bot, listens for termination signals, and gracefully stops the bot.
func (ab *AgoraBot) Run() error {
	errOpen := ab.Session.Open()
	if errOpen != nil {
		return fmt.Errorf("error opening connection: %w", errOpen)
	}

	log.Println("Agora Bot running")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	ab.Session.AddHandler(func(s *discordgo.Session, event *discordgo.Disconnect) {
		log.Println("Connection lost. Reconnecting...")
		for {
			errReconnect := ab.Session.Open()
			if errReconnect == nil {
				log.Println("Reconnected to Discord API")
				break
			}
			log.Printf("Error reconnecting to Discord API: %v. Retrying in 5 seconds...\n", errReconnect)
			time.Sleep(5 * time.Second)
		}
	})

	<-sc

	ab.Session.Close()

	log.Println("Agora Bot stopped")

	return nil
}
