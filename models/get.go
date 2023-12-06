package models

import (
	"github.com/tamiresviegas/warehouse/db"
)

func Get(category string, brand string, maxPrice float64) (produto Produto, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM PRODUCT WHERE  category=$1 AND brand=$2 AND (price BETWEEN 0 AND $3) `, category, brand, maxPrice)

	err = row.Scan(&produto.ID, &produto.Name, &produto.Brand, &produto.Category, &produto.Quantity, &produto.Price)

	return
}

func GetAll() (product []Produto, err error) {

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
		var produto Produto

		err = row.Scan(&produto.ID, &produto.Name, &produto.Brand, &produto.Category, &produto.Quantity, &produto.Price)
		if err != nil {
			continue
		}

		product = append(product, produto)
	}

	return
}
