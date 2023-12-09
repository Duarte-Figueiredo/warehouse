package models

import (
	"github.com/tamiresviegas/warehouse/db"
)

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

		err = row.Scan(&product.Product_ID, &product.Name, &product.Brand, &product.Category, &product.Quantity, &product.Price)
		if err != nil {
			continue
		}

		products = append(products, product)
	}

	return
}

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

		err = row.Scan(&product.Product_ID, &product.Name, &product.Brand, &product.Category, &product.Quantity, &product.Price)
		if err != nil {
			continue
		}

		products = append(products, product)
	}

	return
}
