package response

import "time"

type TransactionResponse struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	ProductID         string    `json:"product_id"`
	OTR               float64   `json:"otr"`
	AdminFee          float64   `json:"admin_fee"`
	InstallmentAmount float64   `json:"installment_amount"`
	AssetName         string    `json:"asset_name"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
}

type TransactionScheduleResponse struct {
	TransactionID string                `json:"transaction_id"`
	Schedule      []PaymentScheduleItem `json:"schedule"`
}

type PaymentScheduleItem struct {
	DueDate time.Time `json:"due_date"`
	Amount  float64   `json:"amount"`
	Status  string    `json:"status"`
}
