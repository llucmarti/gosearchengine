package models

type Product struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Amount    int
	Price     float64
	Materials []Material `gorm:"many2many:product_materials;"`
}

type ProductResponse struct {
	Products []Product `json:"ads"`
	Total    int       `json:"total"`
	Current  int       `json:"current"`
	NextPage int       `json:"nextPage"`
}
