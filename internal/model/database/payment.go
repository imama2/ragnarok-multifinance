package database

import "time"

type Payment struct {
	ID            uint64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TransactionID string     `gorm:"column:transaction_id" json:"transaction_id"`
	DueDate       time.Time  `gorm:"column:due_date" json:"due_date"`
	Amount        float64    `gorm:"column:amount" json:"amount"`
	Status        string     `gorm:"column:status" json:"status"`
	PaidAt        *time.Time `gorm:"column:paid_at" json:"paid_at,omitempty"`
}

func (Payment) TableName() string {
	return "payments"
}
