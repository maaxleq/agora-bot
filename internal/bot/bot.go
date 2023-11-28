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

type AgoraBot struct {
	Conf    config.Config
	Session *discordgo.Session
}

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
