package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var dsn = fmt.Sprintf("%v:%v@/%v", os.Getenv("USERNAME"),
										  os.Getenv("PASSWORD"),
										  os.Getenv("DATABASE"))
func MYSQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
