package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func userReg(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	name := c.DefaultPostForm("name", randUserName())
	password := c.DefaultPostForm("password", "")
	if password == "" {
		c.JSON(200, gin.H{"status": 402})
		return
	}
	sql := `INSERT INTO user (username, password) VALUES (?, ?)`
	rslt, err := db.Exec(sql, name, password)
	if err != nil {
		log.Printf("[ERROR], post userReg : %s", err)
		c.JSON(200, gin.H{"status": 501})
		return
	}

	id, err := rslt.LastInsertId()
	if err != nil {
		log.Printf("[ERROR] get insertid : %s", err)
		c.JSON(200, gin.H{"status": 501})
		return
	}
	user, err := getUserByUid(int(id))
	if err != nil {
		log.Printf("[ERROR] userReg err: %s", err)
	}
	vip := 0
	if user.VipLevel > 0 {
		vip = 1
	}
	c.JSON(200, gin.H{
		"status": 200,
		"data": gin.H{
			"uid":      user.Uid,
			"uname":    user.Uname,
			"stayTime": 0,
			"coin":     user.Coin,
			"score":    user.Score,
			"icon":     user.Icon,
			"exp":      0,
			"lv":       user.VipLevel,
			"pfId":     1,
			"puid":     "10266732",
			"vip":      vip,
		},
		"cookies": []gin.H{
			gin.H{
				"name":   "ghostId",
				"path":   "/",
				"value":  guid(),
				"maxAge": 889032704,
			},
			gin.H{
				"name":   "hallToken",
				"path":   "/",
				"value":  randToken(),
				"maxAge": 7776000,
			},
		},
	})
}

func userRegTest(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"status": 200,
		"data": gin.H{
			"uid":      1,
			"uname":    "tester",
			"stayTime": 0,
			"coin":     10000,
			"score":    0,
			"icon":     4,
			"exp":      0,
			"lv":       1,
			"pfId":     1,
			"puid":     "10266732",
			"vip":      0,
		},
		"cookies": []gin.H{
			gin.H{
				"name":   "ghostId",
				"path":   "/",
				"value":  "c6293367-d56e-4445-a1db-1eb773c2407e",
				"maxAge": 889032704,
			},
			gin.H{
				"name":   "hallToken",
				"path":   "/",
				"value":  "ADMANwBhADcAODExMDUyYzZlNmM3NjMxMjUz",
				"maxAge": 7776000, //90 day
			},
		},
	})
}

/*
 * user not exist = 406
 * no password = 402
 */

func userLogin(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	uid := c.DefaultPostForm("uid", "")
	password := c.DefaultPostForm("password", "")
	if uid == "" || password == "" {
		c.JSON(200, gin.H{"status": 402})
		return
	}
	id, err := strconv.Atoi(uid)
	if err != nil {
		id = 0
	}
	user, status, err := getUserByUidPassword(int(id), password)
	if status == 0 { //user not exist
		c.JSON(200, gin.H{"status": 406})
		return
	} else if status == -1 {
		c.JSON(200, gin.H{
			"message": "账号与密码不符",
			"status":  400,
			"data":    nil,
		})
		return
	}
	vip := 0
	if user.VipLevel > 0 {
		vip = 1
	}
	c.JSON(200, gin.H{
		"status": 200,
		"data": gin.H{
			"uid":      user.Uid,
			"uname":    user.Uname,
			"stayTime": 0,
			"coin":     user.Coin,
			"score":    user.Score,
			"icon":     user.Icon,
			"exp":      0,
			"lv":       user.VipLevel,
			"pfId":     1,
			"puid":     "10266732",
			"vip":      vip,
		},
		"cookies": []gin.H{
			gin.H{
				"name":   "ghostId",
				"path":   "/",
				"value":  guid(),
				"maxAge": 889032704,
			},
			gin.H{
				"name":   "hallToken",
				"path":   "/",
				"value":  randToken(),
				"maxAge": 7776000,
			},
		},
	})
}
