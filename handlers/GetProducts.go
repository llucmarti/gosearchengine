package handlers

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/llucmarti/gosearchengine/database"
	"github.com/llucmarti/gosearchengine/dto"
	"github.com/llucmarti/gosearchengine/helper"
)

func GetProducts(db *database.DB, w http.ResponseWriter, r *http.Request) {

	term := r.URL.Query().Get("term")
	perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
	nPage, _ := strconv.Atoi(r.URL.Query().Get("nPage"))

	products, _ := db.GetProductsByMaterial(term)

	if len(products) == 0 {
		http.Error(w, "No products found for the given material", http.StatusNotFound)
		return
	}

	response := dto.AdResponse{
		Advertising: products,
		Total:       len(products),
		Current:     nPage,
		NextPage:    helper.GetNextPage(nPage, perPage, len(products)),
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
