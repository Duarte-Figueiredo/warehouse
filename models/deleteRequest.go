package models

import "github.com/tamiresviegas/warehouse/db"

func DeleteRequest(reqId int64) (int64, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`delete from request WHERE id = $1`, reqId)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
