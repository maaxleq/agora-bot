package stores

type MongoStore struct {
}

func NewMongoStorer() *MongoStore {
	return &MongoStore{}
}
