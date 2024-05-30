package query

import (
	"github.com/maaxleq/agora-bot/internal/bot"
)

type Query[I interface{}, O interface{}] interface {
	Do(ab *bot.AgoraBot, params I) (O, error)
}
