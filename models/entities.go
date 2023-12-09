package models

type Product struct {
	Product_ID int64   `json:"product_id"`
	Name       string  `json:"name"`
	Brand      string  `json:"brand"`
	Category   string  `json:"category"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
}

// Struct used when message is received from Kafka
type ProductQntUpdt struct {
	Product_ID int64 `json:"product_id"`
	Quantity   int   `json:"quantity"`
}
