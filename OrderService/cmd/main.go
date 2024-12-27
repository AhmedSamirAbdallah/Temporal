package main

import (
	"context"
	"log"
	"net/http"
	"order-service/OrderService/configs"
	"order-service/OrderService/internal/db"
	"order-service/OrderService/internal/handlers"
	"order-service/OrderService/internal/repositories"
	"order-service/OrderService/internal/routes"
	"order-service/OrderService/internal/services"
)

func main() {
	cfg := configs.LoadConfig()
	client, err := db.ConnectDB(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	orderRepo := repositories.NewOrderRepository(client, cfg.DBName, cfg.CollectionName)

	orderService := services.NewOrderService(orderRepo)

	orderHandler := handlers.NewOrderHandler(orderService)

	routes.SetupRoutes(orderHandler)

	// Start the HTTP server on the specified port
	log.Printf("Server started on :%s", cfg.ServerPort)
	if err := http.ListenAndServe(cfg.ServerPort, nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
