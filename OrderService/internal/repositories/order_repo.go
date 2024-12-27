package repositories

import (
	"order-service/OrderService/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// OrderRepository is a repository for managing orders using DBRepository[models.Order].
type OrderRepository struct {
	*DBRepository[models.Order] // Embed the generic repository for orders (inherits all methods and fields from DBRepository)
}

func NewOrderRepository(client *mongo.Client, dbName string, collectionName string) *OrderRepository {
	return &OrderRepository{
		DBRepository: NewDBRepository[models.Order](client, dbName, collectionName),
	}
}
