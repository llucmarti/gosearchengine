package dto

type AdResponse struct {
	Advertising []ProductResponse `json:"ads"`
	Total       int               `json:"total"`
	Current     int               `json:"current"`
	NextPage    int               `json:"nextPage"`
}
