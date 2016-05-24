package main

import "github.com/gin-gonic/gin"

func userReg(c *gin.Context) {
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
				"maxAge": 7776000,
			},
		},
	})
}

func userLogin(c *gin.Context) {
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
				"maxAge": 7776000,
			},
		},
	})
}
