package main

import (
	"context"
	"log"
	"order-service/OrderService/configs"
	"order-service/OrderService/internal/db"
)

func main() {
	cfg := configs.LoadConfig()
	client, err := db.ConnectDB(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)

		}
	}()
}
