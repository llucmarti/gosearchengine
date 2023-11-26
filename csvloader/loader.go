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

func LoadCSV(db *database.DB, filePath string) error {

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

			productErr := db.CreateProduct(&product)
			if productErr != nil {
				log.Printf("Error creating product: %v", productErr)
				// handle error, return or break
			}

			materialErr := db.CreateMaterial(&material)
			if materialErr != nil {
				log.Printf("Error creating material: %v", materialErr)
				// handle error, return or break
			}

			err = db.CreateProductMaterialAssociation(&product, &material)
			if err != nil {
				log.Printf("Error creating association: %v", err)
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
