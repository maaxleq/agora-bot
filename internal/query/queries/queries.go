package queries

import (
	"fmt"

	"github.com/maaxleq/agora-bot/internal/bot"
	"github.com/maaxleq/agora-bot/internal/hub"
	"github.com/maaxleq/agora-bot/internal/store"
)

var empty struct{}

type AddHubQuery struct{}

func (AddHubQuery) Do(ab *bot.AgoraBot, params store.AddHubParams) (struct{}, error) {
	hubsCount, errCount := (*ab.Store).GetHubsCount(store.GetHubsCountParams{})
	if errCount != nil {
		return empty, errCount
	}

	if hubsCount >= ab.Conf.MaxHubs {
		return empty, fmt.Errorf("maximum number of hubs reached")
	}

	err := (*ab.Store).AddHub(params)
	return empty, err
}

type DeleteHubQuery struct{}

func (DeleteHubQuery) Do(ab *bot.AgoraBot, params store.DeleteHubParams) (bool, error) {
	return (*ab.Store).DeleteHub(params)
}

type GetHubQuery struct{}

func (GetHubQuery) Do(ab *bot.AgoraBot, params store.GetHubParams) (hub.Hub, error) {
	return (*ab.Store).GetHub(params)
}

type GetHubsQuery struct{}

func (GetHubsQuery) Do(ab *bot.AgoraBot, params store.GetHubsParams) ([]hub.Hub, error) {
	return (*ab.Store).GetHubs(params)
}

type AddChannelQuery struct{}

func (AddChannelQuery) Do(ab *bot.AgoraBot, params store.AddChannelParams) (struct{}, error) {
	channelsCount, errCount := (*ab.Store).GetChannelsCount(store.GetChannelsCountParams{HubID: params.HubID})
	if errCount != nil {
		return empty, errCount
	}

	if channelsCount >= ab.Conf.MaxChannelsPerHub {
		return empty, fmt.Errorf("maximum number of channels per hub reached")
	}

	err := (*ab.Store).AddChannel(store.AddChannelParams{
		HubID:     params.HubID,
		ChannelID: params.ChannelID,
	})
	return empty, err
}

type DeleteChannelQuery struct{}

func (DeleteChannelQuery) Do(ab *bot.AgoraBot, params store.DeleteChannelParams) (bool, error) {
	return (*ab.Store).DeleteChannel(params)
}

type GetHubsCountQuery struct{}

func (GetHubsCountQuery) Do(ab *bot.AgoraBot, params store.GetHubsCountParams) (uint, error) {
	return (*ab.Store).GetHubsCount(params)
}

type GetChannelsCountQuery struct{}

func (GetChannelsCountQuery) Do(ab *bot.AgoraBot, params store.GetChannelsCountParams) (uint, error) {
	return (*ab.Store).GetChannelsCount(params)
}
