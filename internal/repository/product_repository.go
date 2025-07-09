package repository

import (
	"context"

	"ragnarok-multifinance/internal/model/database"
)

type ProductRepository interface {
	List(ctx context.Context) ([]database.Product, error)
	Create(ctx context.Context, product *database.Product) error
	GetByID(ctx context.Context, id string) (*database.Product, error)
}
