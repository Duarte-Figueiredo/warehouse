package models

type Product struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Brand    string  `json:"brand"`
	Category string  `json:"category"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
