package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Handlers struct {
	MongoClient *mongo.Client
	Context     context.Context
}

func Handler(client *mongo.Client, ctx context.Context) *Handlers {
	return &Handlers{
		MongoClient: client,
		Context:     ctx,
	}
}
