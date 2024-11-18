package stores

import (
	"context"
	"fmt"
	"time"

	"github.com/maaxleq/agora-bot/internal/config"
	"github.com/maaxleq/agora-bot/internal/hub"
	"github.com/maaxleq/agora-bot/internal/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewMongoStorer() *MongoStore {
	return &MongoStore{}
}

func (m *MongoStore) Configure(config config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database to verify connection
	if err = client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	m.client = client
	m.database = client.Database(config.MongoDB)
	m.collection = m.database.Collection("hubs")

	return nil
}

func (m *MongoStore) AddHub(params store.AddHubParams) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if hub already exists
	var existingHub hub.Hub
	err := m.collection.FindOne(ctx, bson.M{"_id": params.Hub.ID}).Decode(&existingHub)
	if err == nil {
		return fmt.Errorf("hub %s already exists", params.Hub.ID.String())
	}
	if err != mongo.ErrNoDocuments {
		return fmt.Errorf("error checking for existing hub: %w", err)
	}

	// Insert the new hub
	_, err = m.collection.InsertOne(ctx, params.Hub)
	if err != nil {
		return fmt.Errorf("failed to insert hub: %w", err)
	}

	return nil
}

func (m *MongoStore) DeleteHub(params store.DeleteHubParams) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := m.collection.DeleteOne(ctx, bson.M{"_id": params.ID})
	if err != nil {
		return false, fmt.Errorf("failed to delete hub: %w", err)
	}

	return result.DeletedCount > 0, nil
}

func (m *MongoStore) GetHub(params store.GetHubParams) (hub.Hub, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result hub.Hub
	err := m.collection.FindOne(ctx, bson.M{"_id": params.ID}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return hub.Hub{}, fmt.Errorf("hub %s not found", params.ID.String())
	}
	if err != nil {
		return hub.Hub{}, fmt.Errorf("failed to get hub: %w", err)
	}

	return result, nil
}

func (m *MongoStore) GetHubs(params store.GetHubsParams) ([]hub.Hub, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get hubs: %w", err)
	}
	defer cursor.Close(ctx)

	var hubs []hub.Hub
	if err = cursor.All(ctx, &hubs); err != nil {
		return nil, fmt.Errorf("failed to decode hubs: %w", err)
	}

	return hubs, nil
}

func (m *MongoStore) AddChannel(params store.AddChannelParams) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := m.collection.UpdateOne(
		ctx,
		bson.M{"_id": params.HubID},
		bson.M{"$addToSet": bson.M{"channels": params.ChannelID}},
	)
	if err != nil {
		return fmt.Errorf("failed to add channel: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("hub %s not found", params.HubID.String())
	}

	return nil
}

func (m *MongoStore) DeleteChannel(params store.DeleteChannelParams) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := m.collection.UpdateOne(
		ctx,
		bson.M{"_id": params.HubID},
		bson.M{"$pull": bson.M{"channels": params.ChannelID}},
	)
	if err != nil {
		return false, fmt.Errorf("failed to delete channel: %w", err)
	}

	return result.ModifiedCount > 0, nil
}

func (m *MongoStore) GetHubsCount(params store.GetHubsCountParams) (uint, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := m.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, fmt.Errorf("failed to get hubs count: %w", err)
	}

	return uint(count), nil
}

func (m *MongoStore) GetChannelsCount(params store.GetChannelsCountParams) (uint, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var hub hub.Hub
	err := m.collection.FindOne(ctx, bson.M{"_id": params.HubID}).Decode(&hub)
	if err == mongo.ErrNoDocuments {
		return 0, fmt.Errorf("hub %s not found", params.HubID.String())
	}
	if err != nil {
		return 0, fmt.Errorf("failed to get hub: %w", err)
	}

	return uint(len(hub.Channels)), nil
}
