package query

import "github.com/maaxleq/agora-bot/internal/store"

type Query[I interface{}, O interface{}] interface {
	Do(s store.Storer, params I) (O, error)
}
