package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/tamiresviegas/warehouse/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	fmt.Println(conf)

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.DataBase)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()

	return conn, err
}
