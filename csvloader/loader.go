package csvloader

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

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
			name := strings.ToLower(strings.ReplaceAll(record[4], " ", ""))

			code := generateCode(name)

			product := models.Product{
				ID:     record[0],
				Name:   record[1],
				Amount: amount,
				Price:  price,
			}

			material := models.Material{
				ID:   code,
				Name: name,
			}

			if err := tx.FirstOrCreate(&material, models.Material{ID: material.ID}).Error; err != nil {
				return err
			}
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

func generateCode(name string) string {
	hash := sha256.Sum256([]byte(name))
	code := hex.EncodeToString(hash[:])[:4]
	return strings.ToUpper(code)
}
