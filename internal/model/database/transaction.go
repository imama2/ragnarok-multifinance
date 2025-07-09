package database

import "time"

type Transaction struct {
	ID                string    `gorm:"column:id;primaryKey;type:char(26)" json:"id"`
	UserID            string    `gorm:"column:user_id" json:"user_id"`
	ProductID         string    `gorm:"column:product_id" json:"product_id"`
	OTR               float64   `gorm:"column:otr" json:"otr"`
	AdminFee          float64   `gorm:"column:admin_fee" json:"admin_fee"`
	InstallmentAmount float64   `gorm:"column:installment_amount" json:"installment_amount"`
	AssetName         string    `gorm:"column:asset_name" json:"asset_name"`
	Status            string    `gorm:"column:status" json:"status"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
}

func (Transaction) TableName() string {
	return "transactions"
}
