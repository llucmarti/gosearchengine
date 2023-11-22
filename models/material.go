package models

type Material struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:product_materials;"`
}
