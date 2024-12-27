package repositories

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// Repository defines the generic interface for CRUD operations on MongoDB entities
type Repository[T any] interface {
	Create(ctx context.Context, record *T) error
	FindByID(ctx context.Context, id string) (*T, error)
	FindAll(ctx context.Context, filter map[string]interface{}) ([]T, error)
	Update(ctx context.Context, id string, update map[string]interface{}) error
	Delete(ctx context.Context, id string) error
}

type DBRepository[T any] struct {
	collection *mongo.Collection
}

func NewDBRepository[T any](client *mongo.Client, dbName string, collectionName string) *DBRepository[T] {
	return &DBRepository[T]{
		collection: client.Database(dbName).Collection(collectionName),
	}
}

func (r *DBRepository[T]) Create(ctx context.Context, record *T) error {
	_, err := r.collection.InsertOne(ctx, record)
	if err != nil {
		log.Println("Failed to create entity:", err)
		return err
	}
	return nil
}

func (r *DBRepository[T]) FindByID(ctx context.Context, id string) (*T, error) {
	var record T
	filter := map[string]interface{}{
		"id": id,
	}
	err := r.collection.FindOne(ctx, filter).Decode(&record)
	if err != nil {
		log.Println("Failed to find entity by ID:", err)
		return nil, err
	}
	return &record, nil
}

func (r *DBRepository[T]) FindAll(ctx context.Context, filter map[string]interface{}) ([]T, error) {
	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		log.Println("Failed to find entities:", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var records []T

	for cur.Next(ctx) {
		var record T
		err := cur.Decode(&record)
		if err != nil {
			log.Println("Failed to decode entity:", err)
			continue
		}
		records = append(records, record)
	}
	return records, nil

}

func (r *DBRepository[T]) Update(ctx context.Context, id string, update map[string]interface{}) error {
	filter := map[string]interface{}{
		"id": id,
	}
	updateDoc := map[string]interface{}{
		"$set": update,
	}
	_, err := r.collection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		log.Println("Failed to update entity:", err)
		return err
	}
	return nil
}

func (r *DBRepository[T]) Delete(ctx context.Context, id string) error {
	filter := map[string]interface{}{
		"id": id,
	}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Failed to delete entity:", err)
		return err
	}
	return nil
}
