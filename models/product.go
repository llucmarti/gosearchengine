package models

type Product struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Amount    int
	Price     float64
	Materials []Material `gorm:"many2many:product_materials;"`
}
