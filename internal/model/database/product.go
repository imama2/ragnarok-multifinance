package database

type Product struct {
	ID          string  `gorm:"column:id;primaryKey;type:char(26)" json:"id"`
	Name        string  `gorm:"column:name" json:"name"`
	Tenure      int     `gorm:"column:tenure" json:"tenure" validate:"oneof=1 2 3 6"`
	MaxLimit    float64 `gorm:"column:max_limit" json:"maxLimit"`
	Interest    float64 `gorm:"column:interest_rate" json:"interest"`
	Description string  `gorm:"column:description" json:"description"`
	IsActive    bool    `gorm:"column:is_active" json:"isActive"`
}

func (Product) TableName() string {
	return "products"
}
