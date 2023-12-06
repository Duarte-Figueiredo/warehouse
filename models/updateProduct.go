package models

import "github.com/tamiresviegas/warehouse/db"

func UpdateProduct(product Product) (int64, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE product SET quantity = $1 WHERE id = $2`, product.Quantity, product.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
