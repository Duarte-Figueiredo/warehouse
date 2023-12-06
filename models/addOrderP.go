package models

import (
	"github.com/tamiresviegas/warehouse/db"
)

func AddOrderP(orderP OrderP) (id int64, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO ORDERP (CategoriesProducts, MaxPrices, Quantities) values ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, orderP.CategoriesProducts, orderP.MaxPrices, orderP.Quantities).Scan(&id)

	return
}
