package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	ws := gin.Default()
	m := melody.New()
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	ws.GET("/speaker", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(func(s *melody.Session) {
		s.Write([]byte("[]"))
	})
	m.HandleDisconnect(func(s *melody.Session) {
		m.BroadcastOthers([]byte("dis "), s)
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastOthers(msg, s)
	})

	go ws.Run(":9030")

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
