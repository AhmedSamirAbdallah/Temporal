package handlers

import (
	"encoding/json"
	"net/http"
	"order-service/OrderService/internal/models"
	"order-service/OrderService/internal/services"
)

type OrderHandler struct {
	OrderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: orderService}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// 1. Parse the request body into an Order object
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}

	// 2. Use the OrderService to create the order
	err = h.OrderService.CreateOrder(&order)
	if err != nil {
		http.Error(w, "Failed to create order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Respond with the created order and a 201 status
	w.WriteHeader(http.StatusCreated)
}
