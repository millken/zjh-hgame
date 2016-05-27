package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DtActionParam struct {
	GameId     int `json:"gameId"`
	RequestId  int `json:"requestId"`
	Action     string
	ReturnType string `json:"returnType"`
	Start      int
	Size       int
	Uid        int
}

func dtSign(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"status": 200,
		"data": gin.H{
			"tkey":       guid(),
			"skey":       guid(),
			"expireIn":   60,
			"expireTime": time.Now().UnixNano() / 1000000,
		},
	})
}

//v=2.0&appId=15&nativeApp=0&ghostId=0419aa7b-eb98-49b0-92a0-379569c2533d&hallToken=ADAAZgBjAGIAZTE4MWYyYzc0Mjk3YTNhNmRs
//&ch=XX_02_001&ptype=页面&tkey=1100353a-a3cd-4dc8-b678-7a0b2d55af3c&sign=579eb4a9c36fc11cae22fb95b885c97c&param=[{"returnType":"payPolicy","requestId":1,"action":"/dt.conf"}]
//验证不过 = 403
func dtAction(c *gin.Context) {
	var p []DtActionParam
	c.Header("Access-Control-Allow-Origin", "*")
	id := strToInt(c.Query("id"))
	v := c.DefaultPostForm("v", "2.0")
	appId := strToInt(c.DefaultPostForm("appId", "15"))
	nativeApp := strToInt(c.DefaultPostForm("nativeApp", "0"))
	ghostId := c.DefaultPostForm("ghostId", "")
	hallToken := c.DefaultPostForm("hallToken", "")
	ch := c.DefaultPostForm("ch", "XX_02_001")
	tkey := c.DefaultPostForm("tkey", "")
	sign := c.DefaultPostForm("sign", "")
	param := c.DefaultPostForm("param", "[]")
	ptype := c.DefaultPostForm("ptype", "页面")

	log.Printf("v: %s, appId: %d, nativeApp: %d, ghostId: %s, hallToken: %s, id: %d, tkey: %s, sign: %s, param: %s, ptype: %s, ch: %s",
		v, appId, nativeApp, ghostId, hallToken, id, tkey, sign, param, ptype, ch)
	if err = json.Unmarshal([]byte(param), &p); err != nil {
		log.Printf("[ERROR] json_decode : %s", err)
		c.JSON(200, gin.H{"status": 403})
		return
	}
	log.Printf("[DEBUG] param : %v", p)
	message := `{}`
	c.String(http.StatusOK, message)
}

func dtLogs(c *gin.Context) {
	t := c.DefaultPostForm("type", "")
	msg := c.DefaultPostForm("msg", "")
	log.Printf("type : %s, msg : %s", t, msg)
}
