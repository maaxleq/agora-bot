// Package bot defines the structure and functionality of the AgoraBot.
package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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
	<-sc

	ab.Session.Close()

	log.Println("Agora Bot stopped")

	return nil
}
