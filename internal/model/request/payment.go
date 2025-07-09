package request

type PaymentCreateRequest struct {
	TransactionID string  `json:"transaction_id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
}
