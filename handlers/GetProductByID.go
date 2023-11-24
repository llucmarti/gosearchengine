package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/llucmarti/gosearchengine/database"
	"github.com/llucmarti/gosearchengine/dto"
	"gorm.io/gorm"
)

func GetProductsByID(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	product, _ := database.GetProductByID(db, id)

	if product.ID == "" {
		http.Error(w, "No products found for this ID", http.StatusNotFound)
		return
	}

	relatedAds, _ := database.GetRelatedProducts(db, id)

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
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
