package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tamiresviegas/warehouse/models"
)

// "Clients should be able to see the received orders"
func GetAllOrders(w http.ResponseWriter, r *http.Request) {

	orders, err := models.GetAllOrders()
	if err != nil {
		log.Printf("Error getting orders: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)

}
