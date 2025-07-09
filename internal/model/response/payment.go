package response

import "time"

type PaymentResponse struct {
	ID            uint64     `json:"id"`
	TransactionID string     `json:"transaction_id"`
	DueDate       time.Time  `json:"due_date"`
	Amount        float64    `json:"amount"`
	Status        string     `json:"status"`
	PaidAt        *time.Time `json:"paid_at,omitempty"`
}
