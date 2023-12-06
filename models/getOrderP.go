package models

import "github.com/tamiresviegas/warehouse/db"

func GetOrderP(id int64) (orderP OrderP, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM ORDERP WHERE id=$1`, id)

	err = row.Scan(&orderP.ID, &orderP.CategoriesProducts, &orderP.MaxPrices, &orderP.Quantities)

	return
}
