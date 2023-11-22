package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/llucmarti/gosearchengine/csvloader"
	"github.com/llucmarti/gosearchengine/db"
	"github.com/llucmarti/gosearchengine/handlers"
)

func main() {
	router := mux.NewRouter()

	db := db.DBconnect()
	fmt.Println(db)

	csvloader.LoadCSV("ad.csv")

	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	http.ListenAndServe(":8080", router)

}
