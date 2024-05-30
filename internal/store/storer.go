package store

import (
	"github.com/maaxleq/agora-bot/internal/config"
	"github.com/maaxleq/agora-bot/internal/hub"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddHubParams struct {
	Hub hub.Hub
}

type DeleteHubParams struct {
	ID primitive.ObjectID
}

type GetHubParams struct {
	ID primitive.ObjectID
}

type GetHubsParams struct{}

type AddChannelParams struct {
	HubID     primitive.ObjectID
	ChannelID string
}

type DeleteChannelParams struct {
	HubID     primitive.ObjectID
	ChannelID string
}

type GetHubsCountParams struct{}

type GetChannelsCountParams struct {
	HubID primitive.ObjectID
}

type Storer interface {
	Configure(config config.Config) error

	AddHub(params AddHubParams) error
	DeleteHub(params DeleteHubParams) (bool, error)
	GetHub(params GetHubParams) (hub.Hub, error)
	GetHubs(params GetHubsParams) ([]hub.Hub, error)
	AddChannel(params AddChannelParams) error
	DeleteChannel(params DeleteChannelParams) (bool, error)
	GetHubsCount(params GetHubsCountParams) (uint, error)
	GetChannelsCount(params GetChannelsCountParams) (uint, error)
}
