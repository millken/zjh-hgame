package gs

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Player struct {
	token string
	uid   int
}

func newPlayer() *Player {
	return &Player{}
}

func (p *Player) ReConn(m map[string]interface{}) (result []gin.H) {
	log.Printf("rrrr %+v", m)
	if m["uid"] != nil && m["token"] != nil {
		p.uid = int(m["uid"].(float64))
		p.token = m["token"].(string)
		log.Printf("[FINE] uid=%d, token=%s", p.uid, p.token)
	}
	result = []gin.H{
		gin.H{
			"code":    0,
			"msgType": 0,
			"b":       1,
			"a":       2,
			"action":  0,
		},
	}
	return
}
