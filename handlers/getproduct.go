package handlers

import (
	"database/sql"
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

	product, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter product: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

// "Clients should be able to get products based on product category, brand and máximum price"
// If one of the filds comes empty, it won't filter by that field.
func Get(w http.ResponseWriter, r *http.Request) {

	category := chi.URLParam(r, "category")
	brand := chi.URLParam(r, "brand")
	maxPrice, errMixPrice := strconv.ParseFloat(chi.URLParam(r, "maxPrice"), 32)
	if errMixPrice != nil {
		log.Printf("Parse Error in maxPrice: %v", errMixPrice)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	product, err := models.Get(category, brand, maxPrice)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("There are no products with the category: %v, brand: %v, and maxPrice: %v", category, brand, maxPrice)
			fmt.Fprintf(w, "There are no products with the category: %s, brand: %s, and maxPrice: %s", category, brand, fmt.Sprintf("%f", maxPrice))
		} else {
			log.Printf("Erro ao atualizar registro: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
