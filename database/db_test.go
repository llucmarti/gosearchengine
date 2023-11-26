package database

import (
	"testing"

	"github.com/llucmarti/gosearchengine/dto"
	"github.com/llucmarti/gosearchengine/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDatastore struct {
	mock.Mock
}

func (m *MockDatastore) CreateMaterial(material *models.Material) error {
	args := m.Called(material)
	return args.Error(0)
}

func TestCreateMaterial(t *testing.T) {
	mockDatastore := new(MockDatastore)
	mockDatastore.On("CreateMaterial", &models.Material{ID: "1", Name: "cobre"}).Return(nil)

	err := mockDatastore.CreateMaterial(&models.Material{ID: "1", Name: "cobre"})

	assert.NoError(t, err)
	mockDatastore.AssertCalled(t, "CreateMaterial", &models.Material{ID: "1", Name: "cobre"})

}

func (m *MockDatastore) CreateProduct(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func TestCreateProduct(t *testing.T) {
	mockDatastore := new(MockDatastore)
	mockDatastore.On("CreateProduct", &models.Product{ID: "1", Name: "cobre", Amount: 10, Price: 6.7}).Return(nil)

	err := mockDatastore.CreateProduct(&models.Product{ID: "1", Name: "cobre", Amount: 10, Price: 6.7})

	assert.NoError(t, err)
	mockDatastore.AssertCalled(t, "CreateProduct", &models.Product{ID: "1", Name: "cobre", Amount: 10, Price: 6.7})

}

func (m *MockDatastore) CreateProductMaterial(product *models.Product, material models.Material) error {
	args := m.Called(product, material)
	return args.Error(0)
}

func TestCreateProductMaterial(t *testing.T) {
	mockDatastore := new(MockDatastore)
	product := &models.Product{ID: "1", Name: "cobre", Amount: 10, Price: 6.7}
	material := models.Material{ID: "12", Name: "lingotes"}
	mockDatastore.On("CreateProductMaterial", product, material).Return(nil)

	err := mockDatastore.CreateProductMaterial(product, material)

	assert.NoError(t, err)
	mockDatastore.AssertCalled(t, "CreateProductMaterial", product, material)
}

func (m *MockDatastore) GetProductsByMaterial(term string) ([]dto.ProductResponse, error) {
	args := m.Called(term)
	return args.Get(0).([]dto.ProductResponse), args.Error(1)
}

func TestGetProductsByMaterial(t *testing.T) {
	mockDatastore := new(MockDatastore)
	mockDatastore.On("GetProductsByMaterial", "cobre").Return([]dto.ProductResponse{{ID: "1", Name: "Product 1"}}, nil)

	products, err := mockDatastore.GetProductsByMaterial("cobre")

	assert.NoError(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, "1", products[0].ID)
	assert.Equal(t, "Product 1", products[0].Name)
}

func (m *MockDatastore) GetProductByID(id string) (*dto.ProductResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.ProductResponse), args.Error(1)
}

func TestGetProductByID(t *testing.T) {
	mockDatastore := new(MockDatastore)
	mockDatastore.On("GetProductByID", "1").Return(&dto.ProductResponse{ID: "1"}, nil)

	product, err := mockDatastore.GetProductByID("1")

	assert.Equal(t, "1", product.ID)
	assert.Nil(t, err)
}
func (m *MockDatastore) GetMaterialsByID(id string) ([]string, error) {
	args := m.Called(id)
	return args.Get(0).([]string), args.Error(1)
}

func TestGetMaterialsByID(t *testing.T) {
	mockDatastore := new(MockDatastore)
	mockDatastore.On("GetMaterialsByID", "1").Return([]string{"cobre", "lingotes"}, nil)

	materials, err := mockDatastore.GetMaterialsByID("1")

	assert.NoError(t, err)
	assert.Equal(t, []string{"cobre", "lingotes"}, materials)
}

func (m *MockDatastore) GetRelatedProducts(id string) ([]dto.ProductResponse, error) {
	args := m.Called(id)
	return args.Get(0).([]dto.ProductResponse), args.Error(1)
}

func TestGetRelatedProducts(t *testing.T) {
	mockDatastore := new(MockDatastore)
	mockDatastore.On("GetRelatedProducts", "1").Return([]dto.ProductResponse{{ID: "2", Name: "Product 2"}}, nil)

	products, err := mockDatastore.GetRelatedProducts("1")

	assert.NoError(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, "2", products[0].ID)
	assert.Equal(t, "Product 2", products[0].Name)
}
