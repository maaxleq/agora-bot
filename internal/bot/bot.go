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
	"github.com/maaxleq/agora-bot/internal/query"
	"github.com/maaxleq/agora-bot/internal/store"
	storeloader "github.com/maaxleq/agora-bot/internal/store/loader"
)

// AgoraBot represents a Discord bot with configuration and session information.
type AgoraBot struct {
	Conf    config.Config
	Session *discordgo.Session
	Store   *store.Storer
}

// NewAgoraBot creates a new instance of AgoraBot with the provided configuration.
func NewAgoraBot(conf config.Config) (*AgoraBot, error) {
	dg, errBot := discordgo.New("Bot " + conf.DiscordToken)
	if errBot != nil {
		return nil, fmt.Errorf("error creating Discord session: %w", errBot)
	}

	store, errStore := storeloader.LoadStore(conf)
	if errStore != nil {
		return nil, fmt.Errorf("error loading store: %w", errStore)
	}

	return &AgoraBot{
		Conf:    conf,
		Session: dg,
		Store:   store,
	}, nil
}

func (ab *AgoraBot) GetQueryDeps() query.QueryDeps {
	return query.QueryDeps{
		Store: ab.Store,
		Conf:  ab.Conf,
	}
}

// Run starts the bot, listens for termination signals, and gracefully stops the bot.
func (ab *AgoraBot) Run() error {
	errOpen := ab.Session.Open()
	if errOpen != nil {
		return fmt.Errorf("error opening connection: %w", errOpen)
	}

	// Add message handler
	ab.Session.AddHandler(ab.handleMessage)

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

// handleMessage processes incoming messages and echoes them to other channels in the same hub
func (ab *AgoraBot) handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if message channel is in any hub
	h, errHub := (*ab.Store).GetHubOfChannel(store.GetHubOfChannelParams{ChannelID: m.ChannelID})
	if errHub != nil {
		log.Printf("Error getting hub of channel: %v\n", errHub)
		return
	}

	// Echo message to other channels in the hub
	for _, targetChannelID := range h.Channels {
		if targetChannelID != m.ChannelID {
			// Create webhook message content
			content := fmt.Sprintf("**%s** (from <#%s>):\n%s",
				m.Author.Username,
				m.ChannelID,
				m.Content,
			)

			_, err := s.ChannelMessageSend(targetChannelID, content)
			if err != nil {
				log.Printf("Error sending message to channel %s: %v\n", targetChannelID, err)
			}
		}
	}
}
