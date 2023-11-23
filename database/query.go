package database

import (
	"fmt"

	"github.com/llucmarti/gosearchengine/dto"
	"gorm.io/gorm"
)

func GetProductsByMaterial(db *gorm.DB, term string) ([]dto.ProductResponse, error) {
	var products []dto.ProductResponse

	err := db.Table("products").
		Select("products.id, products.name, products.amount, products.price").
		Joins("join product_materials on product_materials.product_id = products.id").
		Joins("join materials on product_materials.material_id = materials.id").
		Where("materials.name = ?", term).
		Group("products.id").
		Find(&products).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return products, nil
}

func GetProductByID(db *gorm.DB, id string) (dto.ProductResponse, error) {
	var product dto.ProductResponse

	err := db.Table("products").Select("products.id, products.name, products.amount, products.price").
		Where("products.id = ?", id).
		First(&product).Error

	if err != nil {
		return dto.ProductResponse{}, err
	}

	return product, nil
}

func GetMaterialsByID(db *gorm.DB, id string) ([]string, error) {
	materialIDs := []string{}

	err := db.Table("product_materials").Select("material_id").
		Where("product_id = ?", id).
		Pluck("material_id", &materialIDs).Error

	if err != nil {
		return nil, err
	}

	return materialIDs, nil
}

func GetRelatedProducts(db *gorm.DB, materialIDs []string, id string) ([]dto.ProductResponse, error) {
	var relatedAds []dto.ProductResponse

	err := db.Table("products").Select("products.*").
		Joins("join product_materials on products.id = product_materials.product_id").
		Where("product_materials.material_id IN (?)", materialIDs).
		Where("products.id != ?", id).
		Group("products.id").
		Find(&relatedAds).Error

	if err != nil {
		return nil, err
	}

	return relatedAds, nil
}
