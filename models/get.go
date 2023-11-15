package models

import "github.com/tamiresviegas/warehouse/db"

func Get(id int64) (produto Produto, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM PRODUTOS WHERE id=$1`, id)

	err = row.Scan(&produto.ID, &produto.Name)

	return
}

func GetAll() (produtos []Produto, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row, err := conn.Query(`SELECT * FROM PRODUTOS`)
	if err != nil {
		return
	}

	for row.Next() {
		var produto Produto

		err = row.Scan(&produto.ID, &produto.Name)
		if err != nil {
			continue
		}

		produtos = append(produtos, produto)
	}

	return
}
