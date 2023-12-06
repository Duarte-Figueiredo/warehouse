package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/tamiresviegas/warehouse/models"
)

// "Clients should be able to see a list of available products in the warehouse."
func GetAll(w http.ResponseWriter, r *http.Request) {

	products, err := models.GetAllProducts()
	if err != nil {
		log.Printf("Error getting products: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)

}

// "Clients should be able to get products based on product category, brand and maximum price"
// For a product to show it has to satisfy all the fields (category, brand, and maxPrice)
func GetProductsFiltered(w http.ResponseWriter, r *http.Request) {

	// Read parameters from request
	category := chi.URLParam(r, "category")
	brand := chi.URLParam(r, "brand")
	maxPrice, errMixPrice := strconv.ParseFloat(chi.URLParam(r, "maxPrice"), 32)
	if errMixPrice != nil {
		log.Printf("Parse Error in maxPrice: %v", errMixPrice)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Get Products according to filters
	product, err := models.GetProductFiltered(category, brand, maxPrice)
	if err != nil {
		log.Printf("Error while trying to get products: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if len(product) == 0 {
		log.Printf("There are no products with the category: %v, brand: %v, and maxPrice: %v", category, brand, maxPrice)
		fmt.Fprintf(w, "There are no products with the category: %s, brand: %s, and maxPrice: %s", category, brand, fmt.Sprintf("%f", maxPrice))
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}
}
