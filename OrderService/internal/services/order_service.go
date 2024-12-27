package services

import (
	"context"
	"order-service/OrderService/internal/kafka"
	"order-service/OrderService/internal/models"
	"order-service/OrderService/internal/repositories"

	"github.com/google/uuid"
)

type OrderService struct {
	OrderRepository *repositories.OrderRepository
	Producer        *kafka.KafkaProducer
}

func NewOrderService(orderRepository *repositories.OrderRepository, producer *kafka.KafkaProducer) *OrderService {
	return &OrderService{
		OrderRepository: orderRepository,
		Producer:        producer}

}

func (s *OrderService) CreateOrder(order *models.Order) error {
	order.OrderID = uuid.New().String()
	err := s.OrderRepository.Create(context.Background(), order)
	if err != nil {
		return err
	}
	s.Producer.Publish(context.Background(), "DEV-create-order", order)
	return nil
}
