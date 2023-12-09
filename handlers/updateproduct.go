package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/tamiresviegas/warehouse/models"
)

// Receives an array of products and updates them in the database
func UpdateProducts(message kafka.Message) {
	fmt.Println("receive a message: ", string(message.Value)) // [{"product_id":,"quantity":3},{"product_id":"product2","quantity":2},{"product_id":"product3","quantity":1}]

	var prodQuantites []models.ProductQntUpdt

	if err := json.Unmarshal([]byte(string(message.Value)), &prodQuantites); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	for i := 0; i < len(prodQuantites); i++ {
		// Updates each product quantity
		models.UpdateProduct(prodQuantites[i].Product_ID, prodQuantites[i].Quantity)
	}

}
