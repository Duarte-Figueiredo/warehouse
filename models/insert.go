package models

import (
	"github.com/tamiresviegas/warehouse/db"
)

func Insert(produto Produto) (id int64, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO produtos (name) values ($1) RETURNING id`

	err = conn.QueryRow(sql, produto.Name).Scan(&id)

	return
}
