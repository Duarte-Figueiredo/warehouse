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
// It doesn't returns the products with the lower price. It just returns the products within the maximum price defined
// After the order is handeled it gets deleted from the database
func GetOrderP(w http.ResponseWriter, r *http.Request) {

	// Read parameter from the request
	orderPId, err := strconv.Atoi(chi.URLParam(r, "orderPId"))
	if err != nil {
		log.Printf("Error while parsing order id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Get from the DB the order
	orderP, err := models.GetOrderP(int64(orderPId))
	if err != nil {
		log.Printf("Error while trying to get the order: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	categoriesArray, maxPricesArray, quantitiesArray := strgIntoArrays(orderP)

	// Gets the available and not available products from the DB
	productsAvl, orderPUnavailable, errProd := models.GetProductsAvailability(categoriesArray, maxPricesArray, quantitiesArray)
	if errProd != nil {
		log.Printf("Error while trying to get products availability: %v", errProd)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	productOrder := models.ProductOrderP{
		Produto: productsAvl,
		OrderP:  orderPUnavailable,
	}

	jsonData, err := json.Marshal(productOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Detelets order from DB since it was already handled
	rows, errDel := models.DeleteOrderP(int64(orderPId))
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

func strgIntoArrays(orderP models.OrderP) (categoriesArray []string, maxPricesArray []string, quantitiesArray []string) {
	// Take out all the blank spaces
	var categories = strings.ReplaceAll(orderP.CategoriesProducts, " ", "")
	var maxPrices = strings.ReplaceAll(orderP.MaxPrices, " ", "")
	var quantities = strings.ReplaceAll(orderP.Quantities, " ", "")

	// Put inside of arrays: Split string into an array using comas as the delimiter
	categoriesArray = strings.Split(categories, ",")
	maxPricesArray = strings.Split(maxPrices, ",")
	quantitiesArray = strings.Split(quantities, ",")

	// Validate all arrays have the same length
	if (len(categoriesArray) != len(maxPricesArray)) && (len(maxPricesArray) != len(quantitiesArray)) {
		log.Println("Bad Request. Categories, Max Prices and Quantities have different lengths.")
		return
	}

	return
}
