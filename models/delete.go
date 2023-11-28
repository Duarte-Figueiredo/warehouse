package models

import "github.com/tamiresviegas/warehouse/db"

func Delete(id int64) (int64, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`delete from product WHERE id = $2`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
