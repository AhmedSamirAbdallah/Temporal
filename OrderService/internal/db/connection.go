package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(uri string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
		return nil, err

	}
	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("MongoDB ping failed: %v", err)
		return nil, err
	}

	log.Println("Connected to MongoDB successfully")
	return client, nil
}
