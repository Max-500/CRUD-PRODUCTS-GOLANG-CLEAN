package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseMySQL struct {
	Conn *sql.DB
}

func NewDatabase(user, password, host, dbname string) (*DatabaseMySQL, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &DatabaseMySQL{Conn: db}, nil
}
