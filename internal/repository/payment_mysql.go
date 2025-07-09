package repository

import (
	"context"
	"time"

	"ragnarok-multifinance/internal/model/database"

	"gorm.io/gorm"
)

type PaymentMySQL struct {
	db *gorm.DB
}

func NewPaymentMySQL(db *gorm.DB) *PaymentMySQL {
	return &PaymentMySQL{db: db}
}

func (r *PaymentMySQL) Create(ctx context.Context, payment *database.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

func (r *PaymentMySQL) ListByTransaction(ctx context.Context, transactionID string) ([]database.Payment, error) {
	var payments []database.Payment
	if err := r.db.WithContext(ctx).Where("transaction_id = ?", transactionID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *PaymentMySQL) UpdateStatus(ctx context.Context, id uint64, status string, paidAt *string) error {
	updates := map[string]interface{}{"status": status}
	if paidAt != nil {
		parsed, err := time.Parse(time.RFC3339, *paidAt)
		if err == nil {
			updates["paid_at"] = parsed
		}
	}
	return r.db.WithContext(ctx).Model(&database.Payment{}).Where("id = ?", id).Updates(updates).Error
}
