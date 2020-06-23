package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "root"
	password string = "1sampai8"
	database string = "gorestdb"
)

var dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
func MYSQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
