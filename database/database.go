package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var databaseName = "githubIssuesNotificator"

var client *mongo.Client
var ctx = context.Background()

func Database() *mongo.Database {
	ctx := context.Background()
	databaseUrl := os.Getenv("DATABASE_URL")

	clientOptions := options.Client().ApplyURI(databaseUrl)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatalln(err)
	}

	return client.Database(databaseName)
}

func Close() error {
	return client.Disconnect(ctx)
}
