package configs

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"time"
)

func ConnectMongoDB() *mongo.Client {
	_, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(options.Client().ApplyURI(SetEnvVariables()))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

var Client *mongo.Client = ConnectMongoDB()

func GetCollections(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("").Collection(collectionName)
	return collection
}
