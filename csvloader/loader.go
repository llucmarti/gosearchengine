package csvloader

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

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

			price, err := strconv.ParseFloat(record[3], 64)
			if err != nil {
				return err
			}
			price = price / 100

			amount, err := strconv.Atoi(record[2])
			if err != nil {
				return err
			}

			product := models.Product{
				ID:     record[0],
				Name:   record[1],
				Amount: amount,
				Price:  price,
			}

			fmt.Println(product)

			if err := tx.FirstOrCreate(&product, models.Product{ID: product.ID}).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
