package models

import (
	"github.com/tamiresviegas/warehouse/db"
)

// Updates a specific product
func UpdateProduct(id int64, quantity int) (int64, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE product SET quantity = $1 WHERE product_id = $2`, quantity, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Updates a specific product
func UpdateProductName(productAdd ProductsRespSuppliers) (int64, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE product SET brand = $1, name = $2, quantity = 1, price = $3 WHERE category = $4`, productAdd.Brand, productAdd.Name, productAdd.Price, productAdd.Category)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
