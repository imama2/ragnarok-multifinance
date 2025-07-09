package repository

import (
	"context"

	"ragnarok-multifinance/internal/model/database"

	"gorm.io/gorm"
)

type ProductMySQL struct {
	db *gorm.DB
}

func NewProductMySQL(db *gorm.DB) *ProductMySQL {
	return &ProductMySQL{db: db}
}

func (r *ProductMySQL) List(ctx context.Context) ([]database.Product, error) {
	var products []database.Product
	if err := r.db.WithContext(ctx).Where("is_active = ?", true).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductMySQL) Create(ctx context.Context, product *database.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *ProductMySQL) GetByID(ctx context.Context, id string) (*database.Product, error) {
	var product database.Product
	if err := r.db.WithContext(ctx).Where("id = ? AND is_active = ?", id, true).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
