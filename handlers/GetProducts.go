package handlers

import (
	"fmt"
	"net/http"

	//"strconv"

	"github.com/llucmarti/gosearchengine/models"
	"gorm.io/gorm"
)

func GetProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	term := r.URL.Query().Get("term")
	//perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
	//nPage, _ := strconv.Atoi(r.URL.Query().Get("nPage"))

	products := []models.Product{}

	db.Table("products").Select("products.*").
		Joins("join product_materials on product_materials.product_id = products.id").
		Joins("join materials on product_materials.material_id = materials.id").
		Where("materials.name = ?", term).
		Group("products.id").
		Find(&products)

	fmt.Println(products, len(products))

}
