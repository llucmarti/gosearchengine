package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/llucmarti/gosearchengine/csvloader"
	"github.com/llucmarti/gosearchengine/database"
	"github.com/llucmarti/gosearchengine/handlers"
)

func main() {
	router := mux.NewRouter()

	db := database.DBconnect()
	fmt.Println("Database connected", db)

	csvloader.LoadCSV(db, "ad.csv")
	fmt.Println("CSV loaded")

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) { handlers.GetProducts(db, w, r) }).Methods("GET")
	router.HandleFunc("/detail", func(w http.ResponseWriter, r *http.Request) { handlers.GetProductsByID(db, w, r) }).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
