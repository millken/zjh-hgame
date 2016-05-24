package main

import "github.com/gin-gonic/gin"

func roomUserInfo(c *gin.Context) {
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
	})
}
