package database

import (
	"github.com/llucmarti/gosearchengine/models"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, product models.Product) error {

	err := db.FirstOrCreate(&product, models.Product{ID: product.ID}).Error
	if err != nil {
		return err
	}

	return nil
}

func CreateMaterial(db *gorm.DB, material models.Material) error {

	err := db.FirstOrCreate(&material, models.Material{ID: material.ID}).Error
	if err != nil {
		return err
	}

	return nil
}

func CreateProductMaterial(db *gorm.DB, product models.Product, material models.Material) error {
	err := db.Model(&product).Association("Materials").Append(&material)
	if err != nil {
		return err
	}

	return nil
}
