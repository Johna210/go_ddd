package store

import (
	"coffeeco/internal"
	"context"

	"github.com/google/uuid"
)

type Store struct {
	ID              uuid.UUID
	Location        string
	ProductsForSale []internal.Product
}

type Service struct {
	repo Repository
}

func (s Service) GetStoreSpecificDiscount(ctx context.Context,
	storeID uuid.UUID) (float32, error) {
	dis, err := s.repo.GetStoreDiscount(ctx, storeID)
	if err != nil {
		return 0, err
	}
	return float32(dis), nil
}
