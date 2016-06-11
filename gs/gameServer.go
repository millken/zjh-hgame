package gs

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

var actions = map[int]string{
	1:  "ReConn",
	18: "",
}

type GameServer struct {
	players map[*melody.Session]*Player
}

func NewServer() *GameServer {
	return &GameServer{
		players: make(map[*melody.Session]*Player),
	}
}

func (g *GameServer) Connect(s *melody.Session) {
	log.Printf("[FINE] %s connected gameserver.", s.Request.RemoteAddr)
	p := newPlayer()
	g.players[s] = p
}

func (g *GameServer) Disconnect(s *melody.Session) {
	log.Printf("[FINE] %s disconnected gameserver.", s.Request.RemoteAddr)
	delete(g.players, s)
}

func (g *GameServer) Message(s *melody.Session, msg []byte) {
	var prs []gin.H
	log.Printf("[DEBUG] gsmsg = %s", msg)
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
		aid := int(a["a"].(float64))
		if _, ok := actions[aid]; !ok {
			log.Printf("[ERROR] action [%d] not define", aid)
			continue
		}
		prs = g.Invoke(g.players[s], actions[aid], a)
		log.Printf("+++%+v ", prs)
		jsonBytes, _ := json.Marshal(prs)
		s.Write(jsonBytes)
	}

}

func (g *GameServer) Invoke(any interface{}, name string, args ...interface{}) []gin.H {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	return reflect.ValueOf(any).MethodByName(name).Call(inputs)[0].Interface().([]gin.H)
}
