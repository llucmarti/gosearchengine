package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/llucmarti/gosearchengine/csvloader"
	"github.com/llucmarti/gosearchengine/db"
	"github.com/llucmarti/gosearchengine/handlers"
)

func main() {
	router := mux.NewRouter()

	db := db.DBconnect()
	fmt.Println(db)

	//csvloader.LoadCSV(db, "ad.csv")

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetProducts(db, w, r)
	}).Methods("GET")
	http.ListenAndServe(":8080", router)

}
