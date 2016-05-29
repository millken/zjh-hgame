package main

import (
	"log"

	"gopkg.in/redis.v3"

	"./common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var redisclient *redis.Client
var session *common.Session
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
	//db, err = sqlx.Connect("mysql", "root:123456@tcp(192.168.3.57:3306)/zjh")
	db, err = sqlx.Connect("mysql", "root:password@tcp(127.0.0.1:3306)/zjh")
	if err != nil {
		log.Fatalf("connect mysql server error: %s", err)
	}

	redisclient = redis.NewClient(&redis.Options{
		Addr:     "192.168.0.114:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisclient.Ping().Result()
	if err != nil {
		log.Fatalf("connect redis server error: %s", err)
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
