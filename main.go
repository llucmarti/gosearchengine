package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/llucmarti/gosearchengine/csvloader"
	"github.com/llucmarti/gosearchengine/database"
	"github.com/llucmarti/gosearchengine/handlers"
)

func main() {
	router := mux.NewRouter()

	db := database.DBconnect()
	fmt.Println(db)

	csvloader.LoadCSV(db, "ad.csv")

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) { handlers.GetProducts(db, w, r) }).Methods("GET")
	router.HandleFunc("/detail", func(w http.ResponseWriter, r *http.Request) { handlers.GetProductsByID(db, w, r) }).Methods("GET")
	http.ListenAndServe(":8080", router)

}
