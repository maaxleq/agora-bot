package loader

import (
	"fmt"

	"github.com/maaxleq/agora-bot/internal/config"
	"github.com/maaxleq/agora-bot/internal/store"
	"github.com/maaxleq/agora-bot/internal/store/stores"
)

func LoadStore(config config.Config) (*store.Storer, error) {
	var store store.Storer

	switch config.StoreType {
	case "memory":
		store = &stores.MemoryStore{}
	default:
		return nil, fmt.Errorf("store type %s not supported", config.StoreType)
	}

	err := store.Configure(config)
	if err != nil {
		return nil, err
	}

	return &store, nil
}
