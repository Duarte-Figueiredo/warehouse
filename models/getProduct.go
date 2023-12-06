package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tamiresviegas/warehouse/db"
)

// Queries DB to return only the products with the given fields (it has to match all the filtering fields)
func GetProductFiltered(category string, brand string, maxPrice float64) (products []Product, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row, err := conn.Query(`SELECT * FROM PRODUCT WHERE  category=$1 AND brand=$2 AND (price BETWEEN 0 AND $3) `, category, brand, maxPrice)

	if err != nil {
		return
	}

	for row.Next() {
		var product Product

		err = row.Scan(&product.ID, &product.Name, &product.Brand, &product.Category, &product.Quantity, &product.Price)
		if err != nil {
			continue
		}

		products = append(products, product)
	}

	return
}

// Returns all products stored in the DB
func GetAllProducts() (products []Product, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row, err := conn.Query(`SELECT * FROM PRODUCT`)
	if err != nil {
		return
	}

	for row.Next() {
		var product Product

		err = row.Scan(&product.ID, &product.Name, &product.Brand, &product.Category, &product.Quantity, &product.Price)
		if err != nil {
			continue
		}

		products = append(products, product)
	}

	return
}

// Get products from a desired category
func getProductsCategory(category string) (products []Product, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row, err := conn.Query(`SELECT * FROM PRODUCT WHERE  category=$1 `, category)

	if err != nil {
		return
	}

	for row.Next() {
		var product Product

		err = row.Scan(&product.ID, &product.Name, &product.Brand, &product.Category, &product.Quantity, &product.Price)
		if err != nil {
			continue
		}

		products = append(products, product)
	}

	return
}

// Get products availability, form a category, according to their maximum price and quantity
// It's doesn't return the lower prices between all. It returns the needed amount requested within the max price.
// Example: if there is one that is 2 euros and another 3 euros and max price is 4 euros. If 3 euros is read first from the DB that's going to be the one being retrieved.
func GetProductsAvailability(categoriesArray []string, maxPricesArray []string, quantitiesArray []string) (productsAvl []Product, orderPUnavailable OrderP, err error) {

	var unavailableCategories []string
	var unavailableMaxPrice []string
	var unavailableQuantity []string
	var aggregatedErrors []error

	for i := 0; i < len(categoriesArray); i++ {
		var qntdProdNeeded = 0

		// Slects all products from the desired category
		products, err := getProductsCategory(categoriesArray[i])
		if err != nil {
			fmt.Printf("Error while trying to get the product category: %v", err)
			aggregatedErrors = append(aggregatedErrors, err)
		}

		// Gets the amount of products of the needed category
		qntdProdNeeded, errQnt := strconv.Atoi(quantitiesArray[i])
		if errQnt != nil {
			fmt.Println("Error converting string to integer:", errQnt)
			aggregatedErrors = append(aggregatedErrors, errQnt)
		}

		if len(products) == 0 {
			fmt.Println("No products with the specified category")
		} else {
			maxPrice, errFlt := strconv.ParseFloat(maxPricesArray[i], 64)
			if errFlt != nil {
				fmt.Println("Error converting string to float:", errFlt)
				aggregatedErrors = append(aggregatedErrors, errFlt)
			}

			for j := 0; j < len(products) && qntdProdNeeded != 0; j++ {
				if (products[j].Price <= maxPrice) && (products[j].Quantity > 0) {

					if products[j].Quantity >= qntdProdNeeded {
						// Adds product to the needed amount to the array which contains the available products
						product := products[j]
						product.Quantity = qntdProdNeeded
						productsAvl = append(productsAvl, product)

						// Updates the product quantity in the DB
						product = products[j]
						product.Quantity = products[j].Quantity - qntdProdNeeded
						rows, errUpt := UpdateProduct(product)
						if errUpt != nil {
							fmt.Println("Error while updating product:", errUpt)
							aggregatedErrors = append(aggregatedErrors, errUpt)
						} else {
							fmt.Println("Rows updated: ", rows)
						}

						// Updates the quantity needed for the given product
						qntdProdNeeded = 0
					} else if products[j].Quantity < qntdProdNeeded {

						// Ass the available quantities to the array which contains the available products
						productsAvl = append(productsAvl, products[j])

						// Updates the needed quantity
						qntdProdNeeded -= products[j].Quantity

						// Updtaes the products quantity in the DB
						product := products[j]
						product.Quantity = 0
						rows, errUpt := UpdateProduct(product)
						if errUpt != nil {
							fmt.Println("Error while updating product:", errUpt)
							aggregatedErrors = append(aggregatedErrors, errUpt)
						} else {
							fmt.Println("Rows updated: ", rows)
						}
					}
				}
			}

			// If the needed quantity is still bigger than 0 it means that there wasn't any product of the specified category within the maximum price
			if qntdProdNeeded > 0 {
				unavailableCategories = append(unavailableCategories, categoriesArray[i])
				unavailableMaxPrice = append(unavailableMaxPrice, maxPricesArray[i])
				unavailableQuantity = append(unavailableQuantity, strconv.Itoa(qntdProdNeeded))
			}

		}

	}

	// In the end, if the quantites and categories are not empty than creates the unavailable order to be sent to the user
	// Convert arrays into strings. Each element of the array is separated by ','
	unavailableCategoriesStrg := strings.Join(unavailableCategories, ", ")
	unavailableQuantityStrg := strings.Join(unavailableQuantity, ", ")
	maxPricesArrayStrg := strings.Join(unavailableMaxPrice, ", ")

	orderPUnavailable = OrderP{
		CategoriesProducts: unavailableCategoriesStrg,
		MaxPrices:          maxPricesArrayStrg,
		Quantities:         unavailableQuantityStrg,
	}

	return
}
