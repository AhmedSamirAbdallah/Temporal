package routes

import (
	"net/http"
	"order-service/OrderService/internal/handlers"
)

func SetupRoutes(h *handlers.OrderHandler) {
	http.HandleFunc("POST /v1/api/orders", h.CreateOrder)
}
