package models

import (
	"gorm.io/gorm"
)

type Material struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:product_materials;"`
}
