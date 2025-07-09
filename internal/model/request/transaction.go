package request

type TransactionCreateRequest struct {
	UserID    string  `json:"user_id" binding:"required"`
	ProductID string  `json:"product_id" binding:"required"`
	OTR       float64 `json:"otr" binding:"required,gt=0"`
	AdminFee  float64 `json:"admin_fee" binding:"required,gte=0"`
	AssetName string  `json:"asset_name" binding:"required"`
}
