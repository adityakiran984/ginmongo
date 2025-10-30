package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	// "go.mongodb.org/mongo-driver/v2/x/mongo/driver/mongocrypt/options"
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
