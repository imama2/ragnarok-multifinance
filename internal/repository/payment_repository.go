package repository

import (
	"context"

	"ragnarok-multifinance/internal/model/database"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *database.Payment) error
	ListByTransaction(ctx context.Context, transactionID string) ([]database.Payment, error)
	UpdateStatus(ctx context.Context, id uint64, status string, paidAt *string) error
}
