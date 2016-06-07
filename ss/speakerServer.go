package ss

import (
	"encoding/json"
	"log"

	"github.com/olahol/melody"
)

//[{"reconnReason":1,"uid":10266499,"token":"c56d44cf-d58c-482b-9dd4-ae8dd039eeaa-10266499","appId":15,"speakerType":1,"a":1,"b":1}]

type SpeakerServer struct {
	players map[*melody.Session]*Player
}

func NewServer() *SpeakerServer {
	return &SpeakerServer{
		players: make(map[*melody.Session]*Player),
	}
}

func (g *SpeakerServer) Connect(s *melody.Session) {
	log.Printf("[FINE] %s connected gameserver.", s.Request.RemoteAddr)
	p := newPlayer()
	g.players[s] = p
}

func (g *SpeakerServer) Disconnect(s *melody.Session) {
	log.Printf("[FINE] %s disconnected gameserver.", s.Request.RemoteAddr)
	delete(g.players, s)
}

func (g *SpeakerServer) Message(s *melody.Session, msg []byte) {
	var d []interface{}
	err := json.Unmarshal(msg, &d)
	if err != nil {
		log.Printf("[ERROR] parse json err : %s", err)
	}
	if len(d) == 0 {
		s.Write([]byte("[]"))
		return
	}
	for i := 0; i < len(d); i++ {
		a := d[i].(map[string]interface{})
		if a["uid"] != nil && a["token"] != nil {
			uid := int(a["uid"].(float64))
			token := a["token"].(string)
			log.Printf("[FINE] uid=%d, token=%s", uid, token)
		}
		log.Printf("[DEBUG] %+v", d[i])
	}
	s.Write([]byte(`[{"a":2,"b":1,"reqChatInterval":1000,"openChat":1,"minCoin":200000},{"a":6,"num1":0,"num2":0}]`))
	//[{"a":4,"msgs":[{"msgType":1,"sourceType":1,"sourceName":"海海海","msg":"1钻石100元卖，95元收，加Q群112573524","uid":1188888,"type":1}]}]
}
