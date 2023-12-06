package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/tamiresviegas/warehouse/models"
)

// "Clients should be able to receive products from the warehouse."
// A list of available and not available products is returned
// Quantities in the db is updated
// After the request is handeled it gets deleted from the database
func SendProducts(w http.ResponseWriter, r *http.Request) {

	// Read parameter from request
	reqId, err := strconv.Atoi(chi.URLParam(r, "reqId"))
	if err != nil {
		log.Printf("Error while parsing request id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Get from the DB the request
	request, err := models.GetRequest(int64(reqId))
	if err != nil {
		log.Printf("Error while trying to get the request: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Take out all the blank spaces
	var categories = strings.ReplaceAll(request.CategoriesProducts, " ", "")
	var maxPrices = strings.ReplaceAll(request.MaxPrices, " ", "")
	var quantities = strings.ReplaceAll(request.Quantities, " ", "")

	// Put inside of arrays: Split string into an array using comas as the delimiter
	categoriesArray := strings.Split(categories, ",")
	maxPricesArray := strings.Split(maxPrices, ",")
	quantitiesArray := strings.Split(quantities, ",")

	// Validate all arrays have the same length
	if (len(categoriesArray) != len(maxPricesArray)) && (len(maxPricesArray) != len(quantitiesArray)) {
		log.Println("Request not well performed. Categories, Max Prices and Quantities have different sizes.")
		return
	}

	// Gets the available and not available products from the DB
	productsAvl, requestUnv, errProd := models.GetProductsAvailability(categoriesArray, maxPricesArray, quantitiesArray)
	if errProd != nil {
		log.Printf("Error while trying to get products availability: %v", errProd)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	productRequest := models.ProductRequest{
		Produto: productsAvl,
		Request: requestUnv,
	}

	jsonData, err := json.Marshal(productRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// No final vai ter de eliminar o request da base de dados, isto porque já foi tratado => TODO: NÃO VAI TER DE FAZER ROLLBACK ????
	rows, errDel := models.DeleteRequest(int64(reqId))
	if errDel != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Successfully deleted row: %v", rows)

	// Set the content type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response writer
	w.Write(jsonData)
}
