package query

import (
	"github.com/maaxleq/agora-bot/internal/config"
	"github.com/maaxleq/agora-bot/internal/store"
)

type QueryDeps struct {
	Store *store.Storer
	Conf  config.Config
}

type Query[I interface{}, O interface{}] interface {
	Do(qd QueryDeps, params I) (O, error)
}
