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

	router.HandleFunc("/api/ads", func(w http.ResponseWriter, r *http.Request) { handlers.GetAds(db, w, r) }).Methods("GET")
	router.HandleFunc("/api/details", func(w http.ResponseWriter, r *http.Request) { handlers.GetDetailsByID(db, w, r) }).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
