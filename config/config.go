package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// "context"
// "log"
// "time"

// "go.mongodb.org/mongo-driver/mongo"
// "go.mongodb.org/mongo-driver/mongo/options"

// DatabaseConfig holds the MongoDB database configuration.

type DatabaseConfig struct {
	URI string
	Databse string
	Collections []string
}

func Connect(dbConfig *DatabaseConfig)(*mongo.Database, error) {

clientOptions := options.Client().ApplyURI(dbConfig.URI)
client, err := mongo.NewClient(clientOptions)

if err !=nil {
	log.Fatal(err)
	return nil, err
}

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

defer cancel()

err = client.Connect(ctx)

if err != nil {
	log.Fatal(err)
	return nil, err
}

databse := client.Database(dbConfig.Databse)

return databse, nil

}