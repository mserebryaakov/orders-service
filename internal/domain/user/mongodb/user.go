package mongodb

import (
	"context"
	"errors"
	"fmt"
	"orders-service/internal/controller/http/apperror"
	"orders-service/internal/domain/user"
	"orders-service/pkg/logger"

	repository "orders-service/internal/domain/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
	logger     *logger.Logger
}

func NewUserRepository(database *mongo.Database, collectionName string, logger *logger.Logger) repository.UserRepository {
	return &UserRepository{
		collection: database.Collection(collectionName),
		logger:     logger,
	}
}

// Создание пользователя
func (ur *UserRepository) CreateUser(ctx context.Context, user user.User) (string, error) {
	ur.logger.Debug("create user")
	result, err := ur.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("Create user failed: %v", err)
	}

	ur.logger.Debug("convert InsertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	ur.logger.Trace(user)
	return "", fmt.Errorf("Failed to convert object id to hex; oid: %s", oid)
}

// Получение пользовтеля
func (ur *UserRepository) GetUser(ctx context.Context, user user.UserSignUpDTO) (u user.UserSignUpDTO, err error) {
	filter := bson.M{"username": user.Username, "password": user.Password}

	result := ur.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, apperror.ErrNotFound
		}
		return u, fmt.Errorf("Failed to find one user; error : %v", err)
	}
	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("Failed to decode user from db error : %v", err)
	}

	ur.logger.Info("FindOne user: %s", u)
	return u, nil
}
