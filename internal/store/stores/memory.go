package stores

import (
	"fmt"

	"github.com/maaxleq/agora-bot/internal/config"
	"github.com/maaxleq/agora-bot/internal/hub"
	"github.com/maaxleq/agora-bot/internal/store"
)

type MemoryStore struct {
	hubs []hub.Hub
}

func (m *MemoryStore) Configure(config config.Config) error {
	return nil
}

func (m *MemoryStore) AddHub(params store.AddHubParams) error {
	// Check that the hub doesn't already exist
	for _, h := range m.hubs {
		if h.ID == params.Hub.ID {
			return fmt.Errorf("hub %s already exists", params.Hub.ID.String())
		}
	}

	m.hubs = append(m.hubs, params.Hub)
	return nil
}

func (m *MemoryStore) DeleteHub(params store.DeleteHubParams) (bool, error) {
	for i, h := range m.hubs {
		if h.ID == params.ID {
			m.hubs = append(m.hubs[:i], m.hubs[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (m *MemoryStore) GetHub(params store.GetHubParams) (hub.Hub, error) {
	for _, h := range m.hubs {
		if h.ID == params.ID {
			return h, nil
		}
	}
	return hub.Hub{}, fmt.Errorf("hub %s not found", params.ID.String())
}

func (m *MemoryStore) GetHubs(params store.GetHubsParams) ([]hub.Hub, error) {
	return m.hubs, nil
}

func (m *MemoryStore) AddChannel(params store.AddChannelParams) error {
	for i, h := range m.hubs {
		if h.ID == params.HubID {
			m.hubs[i].Channels = append(m.hubs[i].Channels, params.ChannelID)
			return nil
		}
	}
	return fmt.Errorf("hub %s not found", params.HubID.String())
}

func (m *MemoryStore) DeleteChannel(params store.DeleteChannelParams) (bool, error) {
	for i, h := range m.hubs {
		if h.ID == params.HubID {
			for j, c := range h.Channels {
				if c == params.ChannelID {
					m.hubs[i].Channels = append(h.Channels[:j], h.Channels[j+1:]...)
					return true, nil
				}
			}
		}
	}
	return false, nil
}

func (m *MemoryStore) GetHubsCount(params store.GetHubsCountParams) (uint, error) {
	return uint(len(m.hubs)), nil
}

func (m *MemoryStore) GetChannelsCount(params store.GetChannelsCountParams) (uint, error) {
	for _, h := range m.hubs {
		if h.ID == params.HubID {
			return uint(len(h.Channels)), nil
		}
	}
	return 0, fmt.Errorf("hub %s not found", params.HubID.String())
}

func (m *MemoryStore) GetHubOfChannel(params store.GetHubOfChannelParams) (hub.Hub, error) {
	for _, h := range m.hubs {
		for _, c := range h.Channels {
			if c == params.ChannelID {
				return h, nil
			}
		}
	}
	return hub.Hub{}, fmt.Errorf("channel %s not found", params.ChannelID)
}
