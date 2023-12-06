package models

import "github.com/tamiresviegas/warehouse/db"

func UpdateProduct(produto Produto) (int64, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	println("!!!!! quantity: $1 id: $2 ", produto.Quantity, produto.ID)
	res, err := conn.Exec(`UPDATE product SET quantity = $1 WHERE id = $2`, produto.Quantity, produto.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
