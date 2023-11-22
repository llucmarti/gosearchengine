package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	term := r.URL.Query().Get("term")
	perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
	nPage, _ := strconv.Atoi(r.URL.Query().Get("nPage"))

	fmt.Println(term, perPage, nPage)

}
