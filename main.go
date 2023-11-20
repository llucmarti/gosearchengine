package main

import (
	"fmt"

	"github.com/llucmarti/gosearchengine/db"
)

func main() {
	fmt.Println("Search engine")
	db := db.DBconnect()
	fmt.Println(db)

}
