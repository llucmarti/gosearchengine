package dto

type ProductResponse struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Amount int     `json:"amount"`
	Price  float64 `json:"price"`
}
