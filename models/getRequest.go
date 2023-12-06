package models

import "github.com/tamiresviegas/warehouse/db"

func GetRequest(id int64) (request Request, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM REQUEST WHERE id=$1`, id)

	err = row.Scan(&request.ID, &request.CategoriesProducts, &request.MaxPrices, &request.Quantities)

	return
}
