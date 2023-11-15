package models

import "github.com/tamiresviegas/warehouse/db"

func Update(id int64, produto Produto) (int64, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return 0 , err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE produtos SET name = $1 WHERE id = $2`, produto.Name, produto.ID)
	if err != nil {
		return 0 , err
	}

	return res.RowsAffected()
}
