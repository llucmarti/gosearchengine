package csvloader

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/llucmarti/gosearchengine/database"
	"github.com/llucmarti/gosearchengine/helper"
	"github.com/llucmarti/gosearchengine/models"
	"gorm.io/gorm"
)

func LoadCSV(db *gorm.DB, filePath string) error {

	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		defer file.Close()
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	err = db.Transaction(func(tx *gorm.DB) error {

		_, err := reader.Read()
		if err != nil {
			return err
		}

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

			product := models.Product{
				ID:     record[0],
				Name:   record[1],
				Amount: helper.CalculateAmount(record[2]),
				Price:  helper.CalculatePrice(record[3]),
			}

			material := models.Material{
				ID:   helper.GenerateMaterialID(record[4]),
				Name: helper.NormalizeString(record[4]),
			}

			database.CreateProduct(tx, product)
			database.CreateMaterial(tx, material)
			database.CreateProductMaterial(tx, product, material)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
