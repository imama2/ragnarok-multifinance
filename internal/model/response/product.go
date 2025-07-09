package response

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Tenure      int     `json:"tenure"`
	MaxLimit    float64 `json:"maxLimit"`
	Interest    float64 `json:"interest"`
	Description string  `json:"description"`
	IsActive    bool    `json:"isActive"`
}

type ProductListResponse struct {
	Products []ProductResponse `json:"products"`
}
