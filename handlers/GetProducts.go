package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"

	"github.com/llucmarti/gosearchengine/dto"
	"gorm.io/gorm"
)

func GetProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	term := r.URL.Query().Get("term")
	perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
	nPage, _ := strconv.Atoi(r.URL.Query().Get("nPage"))

	products := []dto.ProductResponse{}

	db.Table("products").Select("products.id, products.name, products.amount, products.price").
		Joins("join product_materials on product_materials.product_id = products.id").
		Joins("join materials on product_materials.material_id = materials.id").
		Where("materials.name = ?", term).
		Group("products.id").
		Find(&products)

	current := nPage
	nextPage := nPage + 1
	if nextPage*perPage >= len(products) {
		nextPage = -1
	}

	response := dto.AdResponse{
		Advertising: products,
		Total:       len(products),
		Current:     current,
		NextPage:    nextPage,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	fmt.Println(products, len(products))

}
