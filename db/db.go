package db

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sqlx.DB
	err error
)

func Boot(cf string) {
	//db, err = sqlx.Connect("mysql", "root:123456@tcp(192.168.3.57:3306)/zjh")
	db, err = sqlx.Connect("mysql", cf)
	if err != nil {
		log.Fatalf("connect mysql server error: %s", err)
	}
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args)
}
