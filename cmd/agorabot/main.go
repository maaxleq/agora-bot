package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/maaxleq/agora-bot/internal/bot"
	"github.com/maaxleq/agora-bot/internal/config"
)

// addHandlers adds given handlers to the bot's session.
func addHandlers(ab *bot.AgoraBot, handlers ...interface{}) {
	for _, handler := range handlers {
		ab.Session.AddHandler(handler)
	}
}

func main() {
	conf, errConf := config.NewFromEnv()
	if errConf != nil {
		log.Fatalf("agorabot: %s", errConf)
	}

	agorabot, errBot := bot.NewAgoraBot(conf)
	if errBot != nil {
		log.Fatalf("agorabot: %s", errBot)
	}

	addHandlers(agorabot, messageCreate)

	errStart := agorabot.Run()
	if errStart != nil {
		log.Fatalf("agorabot: %s", errStart)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
