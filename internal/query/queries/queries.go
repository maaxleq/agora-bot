package queries

import (
	"github.com/maaxleq/agora-bot/internal/hub"
	"github.com/maaxleq/agora-bot/internal/store"
)

var empty struct{}

type AddHubQuery struct{}

func (AddHubQuery) Do(s store.Storer, params store.AddHubParams) (struct{}, error) {
	err := s.AddHub(params)
	return empty, err
}

type DeleteHubQuery struct{}

func (DeleteHubQuery) Do(s store.Storer, params store.DeleteHubParams) (bool, error) {
	return s.DeleteHub(params)
}

type GetHubQuery struct{}

func (GetHubQuery) Do(s store.Storer, params store.GetHubParams) (hub.Hub, error) {
	return s.GetHub(params)
}

type GetHubsQuery struct{}

func (GetHubsQuery) Do(s store.Storer, params store.GetHubsParams) ([]hub.Hub, error) {
	return s.GetHubs(params)
}

type AddChannelQuery struct{}

func (AddChannelQuery) Do(s store.Storer, params store.AddChannelParams) (struct{}, error) {
	err := s.AddChannel(store.AddChannelParams{
		HubID:     params.HubID,
		ChannelID: params.ChannelID,
	})
	return empty, err
}

type DeleteChannelQuery struct{}

func (DeleteChannelQuery) Do(s store.Storer, params store.DeleteChannelParams) (bool, error) {
	return s.DeleteChannel(params)
}

type GetHubsCountQuery struct{}

func (GetHubsCountQuery) Do(s store.Storer, params store.GetHubsCountParams) (int, error) {
	return s.GetHubsCount(params)
}

type GetChannelsCountQuery struct{}

func (GetChannelsCountQuery) Do(s store.Storer, params store.GetChannelsCountParams) (int, error) {
	return s.GetChannelsCount(params)
}
