package csvloader

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func LoadCSV(filePath string) error {

	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read and parse the CSV file
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
	return nil
}
