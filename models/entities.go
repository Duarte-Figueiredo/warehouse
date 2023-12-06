package models

type Produto struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Brand    string  `json:"brand"`
	Category string  `json:"category"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Request struct {
	ID                 int64  `json:"id"`
	CategoriesProducts string `json:"categoriesProducts"`
	MaxPrices          string `json:"maxPrices"`
	Quantities         string `json:"quantities"`
}

type ProductRequest struct {
	Produto []Produto `json:"products"`
	Request Request   `json:"request"`
}
