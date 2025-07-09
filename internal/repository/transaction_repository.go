package repository

import (
	"context"

	"ragnarok-multifinance/internal/model/database"
)

type TransactionRepository interface {
	Create(ctx context.Context, tx *database.Transaction) error
	GetByID(ctx context.Context, id string) (*database.Transaction, error)
	ListByUser(ctx context.Context, userID string) ([]database.Transaction, error)
}
