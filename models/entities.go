package models

type Produto struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Brand    string `json:"brand"`
	Category string `json:"category"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
}
