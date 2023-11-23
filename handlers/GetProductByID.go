package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/llucmarti/gosearchengine/dto"
	"gorm.io/gorm"
)

func GetProductsByID(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	product := dto.ProductResponse{}
	relatedAds := []dto.ProductResponse{}

	materialIDs := []string{}
	db.Table("products").Select("products.id, products.name, products.amount, products.price").
		Where("products.id = ?", id).
		First(&product)

	db.Table("product_materials").Select("material_id").
		Where("product_id = ?", id).
		Pluck("material_id", &materialIDs)

	db.Table("products").Select("products.*").
		Joins("join product_materials on products.id = product_materials.product_id").
		Where("product_materials.material_id IN (?)", materialIDs).
		Group("products.id").
		Find(&relatedAds)

	response := dto.DetailResponse{
		ID:         product.ID,
		Name:       product.Name,
		Amount:     product.Amount,
		Price:      product.Price,
		RelatedAds: relatedAds,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	fmt.Println(product)

}
