package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetProductByID(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Set up the expected query
	rows := sqlmock.NewRows([]string{"ID", "Name", "Amount", "Price"}).
		AddRow("1", "Test Product", 10, 100.0)
	mock.ExpectQuery("^SELECT (.+) FROM products WHERE id = ?$").WithArgs("1").WillReturnRows(rows)

	// Call the function to test
	product, err := GetProductByID(db, "1")
	assert.NoError(t, err)

	// Check the result
	assert.Equal(t, "1", product.ID)
	assert.Equal(t, "Test Product", product.Name)
	assert.Equal(t, 10, product.Amount)
	assert.Equal(t, 100.0, product.Price)
}
