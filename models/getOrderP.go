package models

import "github.com/tamiresviegas/warehouse/db"

func GetAllOrders() (ordersP []OrderP, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row, err := conn.Query(`SELECT * FROM ORDERP`)
	if err != nil {
		return
	}

	for row.Next() {
		var order OrderP

		err = row.Scan(&order.ID, &order.CategoriesProducts, &order.MaxPrices, &order.Quantities)
		if err != nil {
			continue
		}

		ordersP = append(ordersP, order)
	}

	return

}

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
