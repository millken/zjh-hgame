package main

import (
	"github.com/millken/zjh-hgame/common"
	"github.com/millken/zjh-hgame/db"
	"github.com/millken/zjh-hgame/rd"
)

var cf *common.Config
var err error

func initDb() {
	db.Boot(cf.Server.Mysql)

	rd.Boot(cf.Server.Redis)
}
