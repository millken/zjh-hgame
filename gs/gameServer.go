package gs

import (
	"log"

	"github.com/olahol/melody"
)

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
	log.Printf("[DEBUG] gsmsg = %s", msg)

}
