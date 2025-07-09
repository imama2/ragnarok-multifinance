package request

type ProductCreateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Tenure      int     `json:"tenure" binding:"required,oneof=1 2 3 6"`
	MaxLimit    float64 `json:"maxLimit" binding:"required,gt=0"`
	Interest    float64 `json:"interest" binding:"required,gte=0"`
	Description string  `json:"description"`
}
