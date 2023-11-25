package database

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/llucmarti/gosearchengine/dto"
	"github.com/llucmarti/gosearchengine/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func DBconnect() *DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.Product{}, &models.Material{})

	return &DB{db}
}

func (db *DB) GetProductsByMaterial(term string) ([]dto.ProductResponse, error) {
	var products []dto.ProductResponse

	err := db.Table("products").
		Select("products.id, products.name, products.amount, products.price").
		Joins("join product_materials on product_materials.product_id = products.id").
		Joins("join materials on product_materials.material_id = materials.id").
		Where("materials.name = ?", term).
		Group("products.id").
		Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (db *DB) GetProductByID(id string) (*dto.ProductResponse, error) {
	var product dto.ProductResponse

	err := db.Table("products").Select("products.id, products.name, products.amount, products.price").
		Where("products.id = ?", id).
		First(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (db *DB) GetMaterialsByID(id string) ([]string, error) {
	materialIDs := []string{}

	err := db.Table("product_materials").Select("material_id").
		Where("product_id = ?", id).
		Pluck("material_id", &materialIDs).Error

	if err != nil {
		return nil, err
	}

	return materialIDs, nil
}

func (db *DB) GetRelatedProducts(id string) ([]dto.ProductResponse, error) {
	var relatedAds []dto.ProductResponse

	materialIDs, _ := db.GetMaterialsByID(id)
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
