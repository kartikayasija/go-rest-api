package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type db struct{
	Client *mongo.Client
	Ctx context.Context
}
var mongoClient *mongo.Client
var mongoContext context.Context
func ConnectToDb() (*mongo.Client, context.Context) {
	client,err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/rest-api")) 

	if err!=nil{
		log.Fatal(err)
	}

	ctx := context.Background()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connected")
	mongoClient=client
	mongoContext=ctx

	return client, ctx
}

func GetDb() *db {
	return &db{
		Client : mongoClient,
		Ctx: mongoContext,
	}
}
