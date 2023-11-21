package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Name      string
	Amount    int
	Price     int
	Materials []Material `gorm:"many2many:product_materials;"`
}
