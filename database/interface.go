package database

import (
	"github.com/llucmarti/gosearchengine/dto"
	"github.com/llucmarti/gosearchengine/models"
)

type Database interface {
	GetProductsByMaterial(term string) ([]dto.ProductResponse, error)
	GetProductByID(id string) (*dto.ProductResponse, error)
	GetMaterialsByID(id string) ([]string, error)
	GetRelatedProducts(id string) ([]dto.ProductResponse, error)

	CreateProduct(product *models.Product) error
	CreateMaterial(material *models.Material) error
	CreateProductMaterial(product *models.Product, material models.Material) error
}
