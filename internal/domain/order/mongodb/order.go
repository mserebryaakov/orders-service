package mongodb

import (
	"context"
	"errors"
	"fmt"
	"orders-service/internal/domain/order"
	"orders-service/pkg/logger"

	repository "orders-service/internal/domain/order"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
	logger     *logger.Logger
}

func New(database *mongo.Database, collectionName string, logger *logger.Logger) repository.OrderRepository {
	return &OrderRepository{
		collection: database.Collection(collectionName),
		logger:     logger,
	}
}

func (or *OrderRepository) Create(ctx context.Context, order order.Order) (string, error) {
	or.logger.Debug("create order")
	result, err := or.collection.InsertOne(ctx, order)
	if err != nil {
		return "", fmt.Errorf("Create order failed: %v", err)
	}

	or.logger.Debug("covert InsertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	or.logger.Trace(order)
	return "", fmt.Errorf("Failed to convert object id to hex; oid: %s", oid)
}

func (or *OrderRepository) FindOne(ctx context.Context, id string) (o order.Order, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return o, fmt.Errorf("Failed to convert hex to objectid, hex: %s", oid)
	}
	filter := bson.M{"_id": oid}

	result := or.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			// TODO
			//return o, apperror.ErrNotFound
			return o, err
		}
		return o, fmt.Errorf("Failed to find one user by id: %s; error : %v", id, err)
	}
	if err = result.Decode(&o); err != nil {
		return o, fmt.Errorf("Failed to decode order (id: %s) from db error : %v", id, err)
	}

	or.logger.Tracef("FindOne order: %s", o)
	return o, nil
}

func (or *OrderRepository) Update(ctx context.Context, order order.Order) error {
	objectID, err := primitive.ObjectIDFromHex(order.ID)
	if err != nil {
		return fmt.Errorf("Failed to convert user ID to ObjectID, ID: %s", order.ID)
	}

	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(order)
	if err != nil {
		return fmt.Errorf("Failed to marshal order, error: %v", err)
	}

	var updateUserObj bson.M

	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal order bytes, error: %v", err)
	}

	delete(updateUserObj, "_id")

	update := bson.M{
		"$set": updateUserObj,
	}

	result, err := or.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("Failed to execute update order query, error: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("Order not found")
		// TODO
		//return apperror.ErrNotFound
	}

	or.logger.Tracef("Matched %d documents and Modified %d documents", result.MatchedCount, result.ModifiedCount)

	return nil
}

func (or *OrderRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("Failed to convert order ID to ObjectID, ID: %s", id)
	}

	filter := bson.M{"_id": objectID}

	result, err := or.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("Failed to execite query, err: %v", err)
	}
	if result.DeletedCount == 0 {
		// TODO
		//return apperror.ErrNotFound
		return fmt.Errorf("Deleted count = 0")
	}
	or.logger.Tracef("Deleted %d documents", result.DeletedCount)

	return nil
}
