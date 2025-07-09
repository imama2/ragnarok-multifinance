package repository

import (
	"context"

	"ragnarok-multifinance/internal/model/database"

	"gorm.io/gorm"
)

type TransactionMySQL struct {
	db *gorm.DB
}

func NewTransactionMySQL(db *gorm.DB) *TransactionMySQL {
	return &TransactionMySQL{db: db}
}

func (r *TransactionMySQL) Create(ctx context.Context, tx *database.Transaction) error {
	return r.db.WithContext(ctx).Create(tx).Error
}

func (r *TransactionMySQL) GetByID(ctx context.Context, id string) (*database.Transaction, error) {
	var tx database.Transaction
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&tx).Error; err != nil {
		return nil, err
	}
	return &tx, nil
}

func (r *TransactionMySQL) ListByUser(ctx context.Context, userID string) ([]database.Transaction, error) {
	var txs []database.Transaction
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&txs).Error; err != nil {
		return nil, err
	}
	return txs, nil
}
