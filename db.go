package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var err error

var user = User{}

type User struct {
	Uid               uint32
	Uname             string `db:"username"`
	Password          string
	Coin, Score, Icon uint32
	VipLevel          uint32 `db:"vip_level"`
}

func initDb() {
	db, err = sqlx.Connect("mysql", "root:123456@tcp(192.168.3.57:3306)/zjh")
	if err != nil {
		log.Fatalln(err)
	}
}

func getUserByUid(id int) (user User, err error) {
	uid := uint32(id)
	err = db.Get(&user, "SELECT * FROM user WHERE uid=?", uid)
	return
}

func getUserByUidPassword(id int, password string) (user User, status int, err error) {
	user, err = getUserByUid(id)
	if err != nil {
		return
	}
	if user.Password == password {
		status = 1
	} else {
		status = -1
	}
	return
}
