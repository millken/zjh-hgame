package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/millken/zjh-hgame/common"
	"github.com/millken/zjh-hgame/gs"
	"github.com/millken/zjh-hgame/ss"
	"github.com/olahol/melody"
)

func main() {
	var mode string
	c := flag.String("c", "config.toml", "config path")
	flag.Parse()
	cf, err = common.LoadConfig(*c)
	if err != nil {
		log.Fatalln("read config failed, err:", err)
	}
	switch cf.Server.Mode {
	case "release":
		mode = gin.ReleaseMode
	case "debug":
		mode = gin.DebugMode
	case "test":
		mode = gin.TestMode
	default:
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	sss := ss.NewServer()
	speakerServer := gin.Default()
	ssm := melody.New()
	ssm.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	speakerServer.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	speakerServer.GET("/speaker", func(c *gin.Context) {
		ssm.HandleRequest(c.Writer, c.Request)
	})

	ssm.HandleConnect(sss.Connect)
	ssm.HandleDisconnect(sss.Disconnect)
	ssm.HandleMessage(sss.Message)

	go speakerServer.Run(":9030")

	gss := gs.NewServer()
	gameServer := gin.Default()
	gsm := melody.New()
	gsm.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	gameServer.GET("/game", func(c *gin.Context) {
		gsm.HandleRequest(c.Writer, c.Request)
	})

	gsm.HandleConnect(gss.Connect)
	gsm.HandleDisconnect(gss.Disconnect)
	gsm.HandleMessage(gss.Message)

	go gameServer.Run(":8010")

	user := gin.Default()
	user.POST("/user.reg", userReg)
	user.POST("/user.login", userLogin)

	suser := &http.Server{
		Addr:           ":9020",
		Handler:        user,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 10,
	}
	go suser.ListenAndServe()

	dt := gin.Default()
	dt.POST("/dt.sign", dtSign)
	dt.POST("/dt.action", dtAction)
	dt.POST("/dt.logs", dtLogs)

	sdt := &http.Server{
		Addr:           ":6020",
		Handler:        dt,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 10,
	}
	initDb()
	sdt.ListenAndServe()
}
