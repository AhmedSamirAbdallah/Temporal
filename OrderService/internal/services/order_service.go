package services

import (
	"context"
	"order-service/OrderService/internal/models"
	"order-service/OrderService/internal/repositories"
)

type OrderService struct {
	OrderRepository *repositories.OrderRepository
}

func NewOrderService(orderRepository *repositories.OrderRepository) *OrderService {
	return &OrderService{OrderRepository: orderRepository}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	err := s.OrderRepository.Create(context.Background(), order)
	if err != nil {
		return err
	}
	//todo add the produc on the kafka topic

	return nil
}
