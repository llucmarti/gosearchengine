package dto

type DetailResponse struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Amount     int               `json:"amount"`
	Price      float64           `json:"price"`
	RelatedAds []ProductResponse `json:"related_ads"`
}
