package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var ErrNoDiscount = errors.New("no discount for store")

type Repository interface {
	GetStoreDiscount(ctx context.Context, storeID uuid.UUID) (float32, error)
}

type MongoRepository struct {
	storeDiscounts *mongo.Collection
}

func NewMongoRepo(ctx context.Context, connectionString string) (Repository, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	discounts := client.Database("coffeeco").Collection("store_discounts")

	return &MongoRepository{
		storeDiscounts: discounts,
	}, nil
}

func (m MongoRepository) GetStoreDiscount(ctx context.Context, storeID uuid.UUID) (float32, error) {
	var discount float32
	if err := m.storeDiscounts.FindOne(ctx,
		bson.D{{"store_id", storeID.String()}}).Decode(&discount); err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means the query did not match any documents
			return 0, ErrNoDiscount
		}
		return 0, fmt.Errorf("failed to find discount for store: %w", err)
	}
	return discount, nil
}
