package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNextPage(t *testing.T) {
	assert.Equal(t, 1, GetNextPage(0, 10, 20))
	assert.Equal(t, -1, GetNextPage(1, 10, 20))
}

func TestNormalizeString(t *testing.T) {
	assert.Equal(t, "hello", NormalizeString("Hello "))
	assert.Equal(t, "world", NormalizeString(" World"))
}

func TestCalculatePrice(t *testing.T) {
	assert.Equal(t, 1.0, CalculatePrice("100"))
	assert.Equal(t, 2.0, CalculatePrice("200"))
}

func TestCalculateAmount(t *testing.T) {
	assert.Equal(t, 100, CalculateAmount("100"))
	assert.Equal(t, 200, CalculateAmount("200"))
}

func TestGenerateMaterialID(t *testing.T) {
	assert.Equal(t, "044F", GenerateMaterialID(" lingotes"))
	assert.Equal(t, "DC9F", GenerateMaterialID("COBRE"))
}
