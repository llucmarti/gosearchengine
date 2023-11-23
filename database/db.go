package database

import (
	"github.com/llucmarti/gosearchengine/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBconnect() *gorm.DB {
	dsn := "host=localhost user=scrapad password=scrapad dbname=searchengine port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.Product{}, &models.Material{})
	return db
}
