package hub

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hub struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	OwnerID  string             `bson:"owner_id" json:"owner_id"`
	Name     string             `bson:"name" json:"name"`
	Channels []string           `bson:"channels" json:"channels"`
}
