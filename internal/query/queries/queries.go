package queries

import (
	"fmt"

	"github.com/maaxleq/agora-bot/internal/hub"
	"github.com/maaxleq/agora-bot/internal/query"
	"github.com/maaxleq/agora-bot/internal/store"
)

var empty struct{}

type AddHubQuery struct{}

func (AddHubQuery) Do(qd query.QueryDeps, params store.AddHubParams) (struct{}, error) {
	hubsCount, errCount := (*qd.Store).GetHubsCount(store.GetHubsCountParams{})
	if errCount != nil {
		return empty, errCount
	}

	if hubsCount >= qd.Conf.MaxHubs {
		return empty, fmt.Errorf("maximum number of hubs reached")
	}

	err := (*qd.Store).AddHub(params)
	return empty, err
}

type DeleteHubQuery struct{}

func (DeleteHubQuery) Do(qd query.QueryDeps, params store.DeleteHubParams) (bool, error) {
	return (*qd.Store).DeleteHub(params)
}

type GetHubQuery struct{}

func (GetHubQuery) Do(qd query.QueryDeps, params store.GetHubParams) (hub.Hub, error) {
	return (*qd.Store).GetHub(params)
}

type GetHubsQuery struct{}

func (GetHubsQuery) Do(qd query.QueryDeps, params store.GetHubsParams) ([]hub.Hub, error) {
	return (*qd.Store).GetHubs(params)
}

type AddChannelQuery struct{}

func (AddChannelQuery) Do(qd query.QueryDeps, params store.AddChannelParams) (struct{}, error) {
	channelsCount, errCount := (*qd.Store).GetChannelsCount(store.GetChannelsCountParams{HubID: params.HubID})
	if errCount != nil {
		return empty, errCount
	}

	if channelsCount >= qd.Conf.MaxChannelsPerHub {
		return empty, fmt.Errorf("maximum number of channels per hub reached")
	}

	err := (*qd.Store).AddChannel(store.AddChannelParams{
		HubID:     params.HubID,
		ChannelID: params.ChannelID,
	})
	return empty, err
}

type DeleteChannelQuery struct{}

func (DeleteChannelQuery) Do(qd query.QueryDeps, params store.DeleteChannelParams) (bool, error) {
	return (*qd.Store).DeleteChannel(params)
}

type GetHubsCountQuery struct{}

func (GetHubsCountQuery) Do(qd query.QueryDeps, params store.GetHubsCountParams) (uint, error) {
	return (*qd.Store).GetHubsCount(params)
}

type GetChannelsCountQuery struct{}

func (GetChannelsCountQuery) Do(qd query.QueryDeps, params store.GetChannelsCountParams) (uint, error) {
	return (*qd.Store).GetChannelsCount(params)
}
